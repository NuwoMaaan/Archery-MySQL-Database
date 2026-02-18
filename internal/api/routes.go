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

	// SQL Queries are predefined and don't take parameters, therefore no URL parameters for something like /ArcheryClub/get-archer/{id}
	// Not ideal for real API but serves the purpose of demonstrating db interaction via http
	router.Get("/get-archer", ArcherClubHandler.GetArchersByID)
	router.Get("/score-above", ArcherClubHandler.GetArchersScoreAboveValue)
	router.Get("/champion-scores", ArcherClubHandler.GetChampionScoresByEvenetID)
	router.Get("/score-by-date", ArcherClubHandler.GetEndCountOrderTotalByDate)
	router.Get("/score-by-total", ArcherClubHandler.GetEndCountOrderTotalByTotal)
	router.Get("/end-count-total", ArcherClubHandler.GetEndCountTotal)
	router.Get("/score-by-roundnum", ArcherClubHandler.GetEndCountTotalByRoundNum)
	router.Get("/round-category-details", ArcherClubHandler.GetRoundAndCategoryDetails)
	router.Get("/personal-best-by-round", ArcherClubHandler.GetPersonalBestByRound)
	router.Get("/personal-best-all-archers", ArcherClubHandler.GetPersonalBestAllArchers)
}
