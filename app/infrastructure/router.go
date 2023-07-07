package infrastructure

import (
	"fmt"
	"os"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/interfaces"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
	"github.com/gin-gonic/gin"
)

// Dispatch is handle routing
func Dispatch(logger usecases.Logger) {
	positionsController, employeesController := setNeccessaryControllers(logger)

	r := gin.Default()
	r.Use(CorrelationIDGenerator())
	r.Use(ErrorHandler(logger))

	r.GET("/positions", positionsController.Index)
	r.GET("/positions/:id", positionsController.Show)
	r.POST("/positions", positionsController.Store)
	r.DELETE("/positions/:id", positionsController.Destroy)

	r.GET("/position/:position_id/employees", employeesController.Index)
	r.GET("/employees/:id", employeesController.Show)
	r.POST("/employees", employeesController.Store)
	r.DELETE("/employees/:id", employeesController.Destroy)

	r.Run(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
}

func setNeccessaryControllers(logger usecases.Logger) (positionsController *interfaces.PositionController, employeesController *interfaces.EmployeeController) {
	if os.Getenv("DB_DRIVER") == PgxDriver {
		dbHandler, err := NewSQLHandler()
		if err != nil {
			panic(err)
		}
		positionRepository := interfaces.PositionPgRepository{
			SQLHandler: dbHandler,
		}
		positionsController = interfaces.NewPositionController(&positionRepository, logger)
		employeeRepository := interfaces.EmployeePgRepository{
			SQLHandler: dbHandler,
		}
		employeesController = interfaces.NewEmployeeController(&employeeRepository, logger)
	} else if os.Getenv("DB_DRIVER") == MongoDriver {
		dbHandler, err := NewMongoDBHandler()
		if err != nil {
			panic(err)
		}
		positionRepository := interfaces.PositionMongoRepository{
			MongoDBHandler: dbHandler,
		}
		positionsController = interfaces.NewPositionController(&positionRepository, logger)
		employeeRepository := interfaces.EmployeeMongoRepository{
			MongoDBHandler: dbHandler,
		}
		employeesController = interfaces.NewEmployeeController(&employeeRepository, logger)
	} else {
		panic("Invalid DB driver")
	}

	return
}
