package routes

import (
	"nbatracker-backend/internal/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handler *handlers.PlayerHandler) {

	// GET /players + POST /players
	mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetPlayers(w, r)
		case http.MethodPost:
			handler.CreatePlayer(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// GET /players/:id
	mux.HandleFunc("/players/", handler.GetPlayerByID)
}
