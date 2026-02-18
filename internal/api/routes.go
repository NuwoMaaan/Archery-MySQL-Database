package api

import (
	"database/sql"
	"net/http"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/api/handler"
	"github.com/go-chi/chi/v5"
)

func LoadRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/ArcheryClub", func(r chi.Router) {
		loadArcherClubRoutes(r, db)
	})

	return router
}

func loadArcherClubRoutes(router chi.Router, db *sql.DB) {
	ArcherClubHandler := handler.NewArcherHandler(db)

	router.Get("/", ArcherClubHandler.GetArchersByID)
	router.Get("/score-above", ArcherClubHandler.GetArchersScoreAboveValue)
}
