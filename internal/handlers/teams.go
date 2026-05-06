package handlers

import (
	"encoding/json"
	"net/http"

	"nbatracker-backend/internal/models"
	"nbatracker-backend/internal/services"
)

type TeamHandler struct {
	Service *services.TeamService
}

// GET /teams
func (h *TeamHandler) GetTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := h.Service.GetTeams()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if teams == nil {
		teams = []models.Team{}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}
