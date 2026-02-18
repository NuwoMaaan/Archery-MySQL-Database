package handler

import (
	"encoding/json"
	"net/http"

	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/repository"
)

func (h *ArcherHandler) GetArchersByID(w http.ResponseWriter, r *http.Request) {
	result, err := repository.GetArcherByID(h.db)
	if err != nil {
		http.Error(w, "Failed to retrieve archer", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (h *ArcherHandler) GetArchersScoreAboveValue(w http.ResponseWriter, r *http.Request) {
	result, err := repository.GetArcherScoreAboveValue(h.db)
	if err != nil {
		http.Error(w, "Failed to retrieve archer", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}
