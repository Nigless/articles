// Package app configures and runs application.
package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"

	"rest-api/config"
	router "rest-api/internal/controller/http/v1"
	"rest-api/internal/repository"
	"rest-api/internal/service"
	"rest-api/pkg/httpserver"
	"rest-api/pkg/logger"
	"rest-api/pkg/postgres"
	"rest-api/pkg/rest"
)

type App struct {
	http     httpserver.Server
	log      *zerolog.Logger
	database *sqlx.DB
}

func (self *App) Start() {
	if _, err := self.database.Conn(context.Background()); err != nil {
		self.log.Fatal().Msgf("failed to connect postgres: %s", err)
	}

	self.http.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		self.log.
			Info().
			Msgf("app - Run - signal: %s", s.String())
	case err := <-self.http.Notify():
		self.log.
			Error().
			Err(err).
			Msg("app - Start - httpServer.Notify")
	}
}

func (self *App) Stop() {
	self.database.Close()
	self.http.Stop()
}

func New(cfg *config.Config) *App {
	log := logger.New(cfg.Log.Level)

	database, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatal().Msgf("failed to open postgres: %s", err)
	}

	service := service.New(repository.New(database))

	// HTTP Server
	handler := rest.New(log).Handler

	router.New(handler, service)
	httpServer := httpserver.New(handler, httpserver.Config{
		ReadTimeout:     5 * time.Second,
		WriteTimeout:    5 * time.Second,
		ShutdownTimeout: 3 * time.Second,
	})

	return &App{*httpServer, &log, database}
}
