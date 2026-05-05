// internal/routes/routes.go
package routes

import (
    "net/http"
    "nbatracker-backend/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux, h *handlers.PlayerHandler) {
    mux.HandleFunc("/players", h.GetPlayers)
    mux.HandleFunc("/players/", h.GetPlayerByID)
}