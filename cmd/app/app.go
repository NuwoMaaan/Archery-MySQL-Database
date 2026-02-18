package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/api"
)

type App struct {
	router http.Handler
	db     *sql.DB
}

func New(database *sql.DB) *App {
	app := &App{
		router: api.LoadRoutes(database),
		db:     database,
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("Server start failure: %s", err)
	}
	return nil
}
