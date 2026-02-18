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
		http.Error(w, "Failed to retrieve", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (h *ArcherHandler) GetChampionScoresByEvenetID(w http.ResponseWriter, r *http.Request) {
	result, err := repository.GetChampionScoresByEventID(h.db)
	if err != nil {
		http.Error(w, "Failed to retrieve", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (h *ArcherHandler) GetEndCountOrderTotalByDate(w http.ResponseWriter, r *http.Request) {
	result, err := repository.GetEndCountOrderTotalByDate(h.db)
	if err != nil {
		http.Error(w, "Failed to retrieve", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (h *ArcherHandler) GetEndCountOrderTotalByTotal(w http.ResponseWriter, r *http.Request) {
	result, err := repository.GetEndCountOrderTotalByTotal(h.db)
	if err != nil {
		http.Error(w, "Failed to retrieve", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (h *ArcherHandler) GetEndCountTotal(w http.ResponseWriter, r *http.Request) {
	result, err := repository.GetEndCountTotal(h.db)
	if err != nil {
		http.Error(w, "Failed to retrieve", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (h *ArcherHandler) GetEndCountTotalByRoundNum(w http.ResponseWriter, r *http.Request) {
	result, err := repository.GetEndCountOrderTotalByRoundNum(h.db)
	if err != nil {
		http.Error(w, "Failed to retrieve", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (h *ArcherHandler) GetPersonalBestAllArchers(w http.ResponseWriter, r *http.Request) {
	result, err := repository.GetPersonalBestAllArchers(h.db)
	if err != nil {
		http.Error(w, "Failed to retrieve", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (h *ArcherHandler) GetPersonalBestByRound(w http.ResponseWriter, r *http.Request) {
	result, err := repository.GetPersonalBestByRound(h.db)
	if err != nil {
		http.Error(w, "Failed to retrieve", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (h *ArcherHandler) GetRoundAndCategoryDetails(w http.ResponseWriter, r *http.Request) {
	result, err := repository.GetRoundAndCategoryDetails(h.db)
	if err != nil {
		http.Error(w, "Failed to retrieve", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}
