package interfaces

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
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
		handleHttpError(w, pc, err)
	}

	handleHttpResponse(w, positions)
}

// Show return response which contain the specified resource of a Position.
func (pc *PositionController) Show(w http.ResponseWriter, r *http.Request) {
	pc.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	positionID, _ := strconv.Atoi(r.URL.Query().Get("id"))

	position, err := pc.PositionInteractor.Show(positionID)
	
	if err != nil {
		handleHttpError(w, pc, err)
	}

	handleHttpResponse(w, position)
}

func handleHttpError(w http.ResponseWriter, pc *PositionController, err error) {
	pc.Logger.LogError("%s", err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(err)
}

func handleHttpResponse(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
