package interfaces

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
	"github.com/go-chi/chi"
)

// A PositionController belong to the interface layer.
type PositionController struct {
	PositionInteractor usecases.PositionInteractor
	Logger             usecases.Logger
}

// NewPositionController returns the resource of Positions.
func NewPositionController(sqlHandler SQLHandler, logger usecases.Logger) *PositionController {
	return &PositionController{
		PositionInteractor: usecases.PositionInteractor{
			PositionRepository: &PositionRepository{
				SQLHandler: sqlHandler,
			},
		},
		Logger: logger,
	}
}

// Index return response which contain a listing of the resource of Positions.
func (pc *PositionController) Index(w http.ResponseWriter, r *http.Request) {
	pc.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	positions, err := pc.PositionInteractor.Index()

	if err != nil {
		handleHTTPError(w, pc.Logger, err)
	}

	handleHTTPResponse(w, positions)
}

// Store is store a newly created resource in storage.
func (pc *PositionController) Store(w http.ResponseWriter, r *http.Request) {
	pc.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	p := domain.Position{}
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		handleHTTPError(w, pc.Logger, err)
	}

	err = pc.PositionInteractor.Store(p)

	if err != nil {
		handleHTTPError(w, pc.Logger, err)
	}

	http.Redirect(w, r, "/positions", http.StatusSeeOther)
}

// Show return response which contain the specified resource of a Position.
func (pc *PositionController) Show(w http.ResponseWriter, r *http.Request) {
	pc.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	positionID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	position, err := pc.PositionInteractor.Show(positionID)

	if err != nil {
		handleHTTPError(w, pc.Logger, err)
	}

	handleHTTPResponse(w, position)
}

// Destroy is remove the specified resource from storage.
func (pc *PositionController) Destroy(w http.ResponseWriter, r *http.Request) {
	pc.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	positionID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := pc.PositionInteractor.Destroy(positionID)

	if err != nil {
		handleHTTPError(w, pc.Logger, err)
	}

	http.Redirect(w, r, "/positions", http.StatusSeeOther)
}
