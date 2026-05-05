package handlers

import (
	"encoding/json"
	"nbatracker-backend/internal/services"
	"net/http"
	"strconv"
	"strings"
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

func (h *PlayerHandler) GetPlayerByID(w http.ResponseWriter, r *http.Request) {
	// URL ejemplo: /players/1
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {
		http.Error(w, "ID requerido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	player, err := h.Service.GetPlayerByID(id)
	if err != nil {
		http.Error(w, "Jugador no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}
