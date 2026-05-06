package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"nbatracker-backend/internal/models"
	"nbatracker-backend/internal/services"
)

type PlayerHandler struct {
	Service *services.PlayerService
}

type paginatedResponse struct {
	Data  []models.Player `json:"data"`
	Total int             `json:"total"`
	Page  int             `json:"page"`
	Limit int             `json:"limit"`
	Pages int             `json:"pages"`
}

// GET /players?page=1&limit=10
func (h *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	page := 1
	limit := 9

	if p := r.URL.Query().Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	if l := r.URL.Query().Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 && v <= 100 {
			limit = v
		}
	}

	players, total, err := h.Service.GetPlayersPaginated(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pages := (total + limit - 1) / limit

	resp := paginatedResponse{
		Data:  players,
		Total: total,
		Page:  page,
		Limit: limit,
		Pages: pages,
	}
	if resp.Data == nil {
		resp.Data = []models.Player{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GET /players/:id
func (h *PlayerHandler) GetPlayerByID(w http.ResponseWriter, r *http.Request) {
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

// POST /players
func (h *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var p models.Player

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// validación de numeros negativos
	if p.Championships < 0 || p.MVP < 0 || p.FinalsMVP < 0 || p.DPOY < 0 || p.ROTY < 0 {
		http.Error(w, "Los valores no pueden ser negativos", http.StatusBadRequest)
		return
	}

	err = h.Service.CreatePlayer(&p)
	if err != nil {
		http.Error(w, "Error al crear jugador", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

// PUT /players/:id
func (h *PlayerHandler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
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

	var p models.Player
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// validacion de campos numericos no negativos
	if p.Championships < 0 || p.MVP < 0 || p.FinalsMVP < 0 || p.DPOY < 0 || p.ROTY < 0 {
		http.Error(w, "Los valores no pueden ser negativos", http.StatusBadRequest)
		return
	}

	err = h.Service.UpdatePlayer(id, &p)
	if err != nil {
		http.Error(w, "Error al actualizar jugador", http.StatusInternalServerError)
		return
	}

	// devolver jugador actualizado
	updatedPlayer, err := h.Service.GetPlayerByID(id)
	if err != nil {
		http.Error(w, "Error al obtener jugador actualizado", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPlayer)
}

// DELETE /players/:id
func (h *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
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

	err = h.Service.DeletePlayer(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Jugador no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al eliminar jugador", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}