package main

import (
	"UEBA/config"
	"UEBA/internal/csvreader"
	"UEBA/pkg/zaplogger"
	"UEBA/transport/http"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Viper
	_, cfg, errViper := config.NewViper("conf_local")
	if errViper != nil {
		log.Fatal(errors.WithMessage(errViper, "Viper startup error"))
	}

	// Zap logger
	logger, loggerCleanup, errZapLogger := zaplogger.New(zaplogger.Mode(cfg.Logger.Development))
	if errZapLogger != nil {
		log.Fatal(errors.WithMessage(errZapLogger, "Zap logger startup error"))
	}

	// CSVReader
	csvreader := csvreader.NewCsvReader(cfg, logger)

	// HTTP Controller
	uebaController := http.NewConfigController(logger, csvreader)

	// HTTP router
	router := http.NewRouter(cfg, logger, uebaController)
	router.RegisterRoutes()

	// Channel for error transmission
	errCh := make(chan error, 1)

	// Router in goroutine
	go func() {
		err := router.Start()
		if err != nil {
			logger.Error("Error starting router", zap.Error(err))
		}
		errCh <- err
	}()

	// Handle shutdown gracefully
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		logger.Error(err.Error())
	case <-shutdown:
	}
	logger.Info("Received shutdown signal. Shutting down...")
	loggerCleanup()
	logger.Info("Application stopped gracefully")

}
