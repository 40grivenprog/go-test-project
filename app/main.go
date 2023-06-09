package main

import (
	"os"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/infrastructure"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Pgx is postgres driver name
const Pgx string = "pgx"

// Mongo is mongo driver name
const Mongo string = "mongo"

func main() {
	logger := infrastructure.NewLogger()

	infrastructure.Load(logger)

	var dbHandler interface{}
	var err error

	if os.Getenv("DB_DRIVER") == Pgx {
		dbHandler, err = infrastructure.NewSQLHandler()
	} else if os.Getenv("DB_DRIVER") == Mongo {
		dbHandler, err = infrastructure.NewMongoDBHandler()
	} else {
		panic("Invalid db driver")
	}

	if err != nil {
		logger.LogError("%s", err)
	}

	infrastructure.Dispatch(logger, dbHandler)
}
