package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const appVersion = "3.0.0"

type serverConfig struct {
	port        int
	environment string
}

type applicationDependences struct {
	config serverConfig
	logger *slog.Logger
}

func main() {
	var settings serverConfig
	flag.IntVar(&settings.port, "port", 4000, "Server Port")
	flag.StringVar(&settings.environment, "env", "development", "Environment(development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	appInstance := &applicationDependences{
		config: settings,
		logger: logger,
	}

	apiServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", settings.port),
		Handler:      appInstance.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}
	logger.Info("Starting Server", "address", apiServer.Addr, "environment", settings.environment)
	err := apiServer.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
