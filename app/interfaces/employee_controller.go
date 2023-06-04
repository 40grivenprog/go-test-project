package interfaces

import (
	"net/http"

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
func NewEmployeeController(dbHandler interface{}, logger usecases.Logger) *EmployeeController {
	var employeeRepository EmployeeRepository

	switch dbHandler.(type) {
	case SQLHandler:
		sqlHandler, _ := dbHandler.(SQLHandler)
		employeeRepository = &EmployeePgRepository{
			SQLHandler: sqlHandler,
		}
	case MongoDBHandler:
		mongoDbHandler, _ := dbHandler.(MongoDBHandler)
		employeeRepository = &EmployeeMongoRepository{
			MongoDBHandler: mongoDbHandler,
		}
	}

	return &EmployeeController{
		EmployeeInteractor: usecases.EmployeeInteractor{
			EmployeeRepository: employeeRepository,
		},
		Logger: logger,
	}
}

// Index is display a listing of the resource.
func (ec *EmployeeController) Index(c *gin.Context) {
	employees, err := ec.EmployeeInteractor.Index(c.Param("position_id"))

	if err != nil {
		c.AbortWithError(CalculateResponseErrorStatus(err), err)
		return
	}

	c.IndentedJSON(http.StatusOK, employees)
}

// Store is stora a newly created resource in storage.
func (ec *EmployeeController) Store(c *gin.Context) {
	employee := domain.Employee{}

	if err := c.BindJSON(&employee); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := ec.EmployeeInteractor.Store(employee)

	if err != nil {
		c.AbortWithError(CalculateResponseErrorStatus(err), err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/positions")
}

// Show return response which contain the specified resource of a employee
func (ec *EmployeeController) Show(c *gin.Context) {
	employee, err := ec.EmployeeInteractor.Show(c.Param("id"))

	if err != nil {
		c.AbortWithError(CalculateResponseErrorStatus(err), err)
		return
	}

	c.IndentedJSON(http.StatusOK, employee)
}

// Destroy is remove the specified resource from storage.
func (ec *EmployeeController) Destroy(c *gin.Context) {
	err := ec.EmployeeInteractor.Destroy(c.Param("id"))

	if err != nil {
		c.AbortWithError(CalculateResponseErrorStatus(err), err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/positions")
}
