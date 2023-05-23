package interfaces

import (
	"net/http"
	"strconv"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
	"github.com/gin-gonic/gin"
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
func (ec *EmployeeController) Index(c *gin.Context) {
	positionID, _ := strconv.Atoi(c.Param("position_id"))

	employees, err := ec.EmployeeInteractor.Index(positionID)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	c.IndentedJSON(http.StatusOK, employees)
}

// Store is stora a newly created resource in storage.
func (ec *EmployeeController) Store(c *gin.Context) {
	employee := domain.Employee{}

	if err := c.BindJSON(&employee); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	err := ec.EmployeeInteractor.Store(employee)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	c.Redirect(http.StatusSeeOther, "/positions")
}

// Show return response which contain the specified resource of a employee
func (ec *EmployeeController) Show(c *gin.Context) {

	employeeID, _ := strconv.Atoi(c.Param("id"))

	employee, err := ec.EmployeeInteractor.Show(employeeID)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	c.IndentedJSON(http.StatusOK, employee)
}

// Destroy is remove the specified resource from storage.
func (ec *EmployeeController) Destroy(c *gin.Context) {
	employeeID, _ := strconv.Atoi(c.Param("id"))

	err := ec.EmployeeInteractor.Destroy(employeeID)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	c.Redirect(http.StatusSeeOther, "/positions")
}
