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

// helper para errores JSON
func jsonError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"error": msg,
	})
}

// GET /players?page=&limit=&q=&sort=&order=
// GetPlayers godoc
// @Summary Listar jugadores
// @Description Obtener lista de jugadores con paginación, búsqueda y ordenamiento
// @Tags players
// @Accept json
// @Produce json
// @Param page query int false "Página"
// @Param limit query int false "Límite"
// @Param q query string false "Búsqueda"
// @Param sort query string false "Campo"
// @Param order query string false "asc|desc"
// @Success 200 {object} paginatedResponse
// @Router /players [get]
func (h *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	page := 1
	limit := 9

	q := r.URL.Query().Get("q")
	sort := r.URL.Query().Get("sort")
	order := r.URL.Query().Get("order")

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

	players, total, err := h.Service.GetPlayersAdvanced(page, limit, q, sort, order)
	if err != nil {
		jsonError(w, err.Error(), http.StatusInternalServerError)
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
// GetPlayerByID godoc
// @Summary Obtener jugador por ID
// @Tags players
// @Param id path int true "ID"
// @Success 200 {object} models.Player
// @Failure 404 {object} map[string]string
// @Router /players/{id} [get]
func (h *PlayerHandler) GetPlayerByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {
		jsonError(w, "ID requerido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		jsonError(w, "ID inválido", http.StatusBadRequest)
		return
	}

	player, err := h.Service.GetPlayerByID(id)
	if err != nil {
		jsonError(w, "Jugador no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}

// POST /players
// CreatePlayer godoc
// @Summary Crear jugador
// @Tags players
// @Accept json
// @Produce json
// @Param player body models.Player true "Jugador"
// @Success 201 {object} models.Player
// @Router /players [post]
func (h *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var p models.Player

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		jsonError(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if p.Championships < 0 || p.MVP < 0 || p.FinalsMVP < 0 || p.DPOY < 0 || p.ROTY < 0 {
		jsonError(w, "Los valores no pueden ser negativos", http.StatusBadRequest)
		return
	}

	err = h.Service.CreatePlayer(&p)
	if err != nil {
		jsonError(w, "Error al crear jugador", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

// PUT /players/:id
// UpdatePlayer godoc
// @Summary Actualizar jugador
// @Tags players
// @Param id path int true "ID"
// @Param player body models.Player true "Jugador"
// @Success 200 {object} models.Player
// @Router /players/{id} [put]
func (h *PlayerHandler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {
		jsonError(w, "ID requerido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		jsonError(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var p models.Player
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		jsonError(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if p.Championships < 0 || p.MVP < 0 || p.FinalsMVP < 0 || p.DPOY < 0 || p.ROTY < 0 {
		jsonError(w, "Los valores no pueden ser negativos", http.StatusBadRequest)
		return
	}

	err = h.Service.UpdatePlayer(id, &p)
	if err != nil {
		jsonError(w, "Error al actualizar jugador", http.StatusInternalServerError)
		return
	}

	updatedPlayer, err := h.Service.GetPlayerByID(id)
	if err != nil {
		jsonError(w, "Error al obtener jugador actualizado", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPlayer)
}

// DELETE /players/:id
// DeletePlayer godoc
// @Summary Eliminar jugador
// @Tags players
// @Param id path int true "ID"
// @Success 204
// @Router /players/{id} [delete]
func (h *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {
		jsonError(w, "ID requerido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		jsonError(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.Service.DeletePlayer(id)
	if err != nil {
		if err == sql.ErrNoRows {
			jsonError(w, "Jugador no encontrado", http.StatusNotFound)
			return
		}
		jsonError(w, "Error al eliminar jugador", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
