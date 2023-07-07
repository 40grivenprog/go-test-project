package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

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


func createCtrlCContext() context.Context {
  ctx, cancel := context.WithCancel(context.Background())
  go func() {
    sigChan := make(chan os.Signal, 1)

    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
    
		<-sigChan
    
		cancel()
  }()

  return ctx
}
