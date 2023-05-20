package interfaces

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
	"github.com/go-chi/chi"
)

// A EmployeeController belong to the interface layer.
type EmployeeController struct {
	EmployeeInteractor usecases.EmployeeInteractor
	Logger             usecases.Logger
}

// NewEmployeeController returns the resource of Employees.
func NewEmployeeController(sqlHandler SQLHandler, logger usecases.Logger) *EmployeeController {
	return &EmployeeController{
		EmployeeInteractor: usecases.EmployeeInteractor{
			EmployeeRepository: &EmployeeRepository{
				SQLHandler: sqlHandler,
			},
		},
		Logger: logger,
	}
}

// Index is display a listing of the resource.
func (ec *EmployeeController) Index(w http.ResponseWriter, r *http.Request) {
	ec.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	positionID, _ := strconv.Atoi(chi.URLParam(r, "position_id"))

	employees, err := ec.EmployeeInteractor.Index(positionID)

	if err != nil {
		handleHTTPError(w, ec.Logger, err)
	}

	handleHTTPResponse(w, employees)
}

// Store is stora a newly created resource in storage.
func (ec *EmployeeController) Store(w http.ResponseWriter, r *http.Request) {
	ec.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	employee := domain.Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		handleHTTPError(w, ec.Logger, err)
	}

	err = ec.EmployeeInteractor.Store(employee)

	if err != nil {
		handleHTTPError(w, ec.Logger, err)
	}

	http.Redirect(w, r, "/positions", http.StatusSeeOther)
}

// Show return response which contain the specified resource of a employee
func (ec *EmployeeController) Show(w http.ResponseWriter, r *http.Request) {
	ec.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	employeeID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	employee, err := ec.EmployeeInteractor.Show(employeeID)

	if err != nil {
		handleHTTPError(w, ec.Logger, err)
	}

	handleHTTPResponse(w, employee)
}

// Destroy is remove the specified resource from storage.
func (ec *EmployeeController) Destroy(w http.ResponseWriter, r *http.Request) {
	ec.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	employeeID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := ec.EmployeeInteractor.Destroy(employeeID)

	if err != nil {
		handleHTTPError(w, ec.Logger, err)
	}

	http.Redirect(w, r, "/positions", http.StatusSeeOther)
}