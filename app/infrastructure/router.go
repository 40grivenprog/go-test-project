package infrastructure

import (
	"net/http"
	"os"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/interfaces"
	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/usecases"
	"github.com/go-chi/chi"
)

// Dispatch is handle routing
func Dispatch(logger usecases.Logger, sqlHandler interfaces.SQLHandler) {
	positionController := interfaces.NewPositionController(sqlHandler, logger)


	r := chi.NewRouter()
	r.Get("/positions", positionController.Index)
	r.Get("/position", positionController.Show)
	r.Post("/positions", positionController.Store)

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r); err != nil {
		logger.LogError("%s", err)
	}
}
