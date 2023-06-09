package infrastructure

import (
	"fmt"
	"os"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/interfaces"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
	"github.com/gin-gonic/gin"
)

// Dispatch is handle routing
func Dispatch(logger usecases.Logger, dbHandler interface{}) {
	positionController := interfaces.NewPositionController(dbHandler, logger)
	employeesController := interfaces.NewEmployeeController(dbHandler, logger)
	usersController := interfaces.NewUserController(dbHandler, logger)

	r := gin.Default()

	public := r.Group("/api")
	protected := r.Group("/api/admin")

	public.Use(CorrelationIDGenerator())
	public.Use(ErrorHandler(logger))
	protected.Use(CorrelationIDGenerator())
	protected.Use(ErrorHandler(logger))
	protected.Use(JwtAuthMiddleware())

	public.POST("/register", usersController.Register)
	public.POST("/login", usersController.Login)

	protected.GET("/positions", positionController.Index)
	protected.GET("/positions/:id", positionController.Show)
	protected.POST("/positions", positionController.Store)
	protected.DELETE("/positions/:id", positionController.Destroy)

	protected.GET("/position/:position_id/employees", employeesController.Index)
	protected.GET("/employees/:id", employeesController.Show)
	protected.POST("/employees", employeesController.Store)
	protected.DELETE("/employees/:id", employeesController.Destroy)

	r.Run(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
}
