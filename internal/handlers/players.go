// internal/handlers/players.go
package handlers

import (
	"encoding/json"
	"nbatracker-backend/internal/services"
	"net/http"
)

type PlayerHandler struct {
	Service *services.PlayerService
}

func (h *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "ok",
	}

	json.NewEncoder(w).Encode(response)
}
