package handlers

import (
	"encoding/json"
	"net/http"

	"nbatracker-backend/internal/services"
)

type PlayerHandler struct {
	Service *services.PlayerService
}

func (h *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := h.Service.GetPlayers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}