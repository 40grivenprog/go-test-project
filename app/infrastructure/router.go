package infrastructure

import (
	"fmt"
	"os"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/interfaces"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
	"github.com/gin-gonic/gin"
)

// Dispatch is handle routing
func Dispatch(logger usecases.Logger, dbHandler interface {}) {
	positionController := interfaces.NewPositionController(dbHandler, logger)
	employeesController := interfaces.NewEmployeeController(dbHandler, logger)

	r := gin.Default()
	r.Use(CorrelationIDGenerator())
	r.Use(ErrorHandler(logger))

	r.GET("/positions", positionController.Index)
	r.GET("/positions/:id", positionController.Show)
	r.POST("/positions", positionController.Store)
	r.DELETE("/positions/:id", positionController.Destroy)

	r.GET("/position/:position_id/employees", employeesController.Index)
	r.GET("/employees/:id", employeesController.Show)
	r.POST("/employees", employeesController.Store)
	r.DELETE("/employees/:id", employeesController.Destroy)

	r.Run(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
}
