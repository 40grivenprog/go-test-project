package main

import (
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

	var err error

	if err != nil {
		logger.LogError("%s", err)
	}

	infrastructure.Dispatch(logger)
}
