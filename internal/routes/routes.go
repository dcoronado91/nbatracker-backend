package routes

import (
	"nbatracker-backend/internal/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, playerHandler *handlers.PlayerHandler, teamHandler *handlers.TeamHandler) {

	mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			playerHandler.GetPlayers(w, r)
		case http.MethodPost:
			playerHandler.CreatePlayer(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/players/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			playerHandler.GetPlayerByID(w, r)
		case http.MethodPut:
			playerHandler.UpdatePlayer(w, r)
		case http.MethodDelete:
			playerHandler.DeletePlayer(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/teams", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			teamHandler.GetTeams(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
}
