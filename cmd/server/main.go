package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	_ "nbatracker-backend/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"nbatracker-backend/internal/db"
	"nbatracker-backend/internal/handlers"
	"nbatracker-backend/internal/repository"
	"nbatracker-backend/internal/routes"
	"nbatracker-backend/internal/services"
)

// @title NBA Tracker API
// @version 1.0
// @description API para manejar jugadores NBA
// @host localhost:8080
// @BasePath /

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar .env, usando variables del sistema")
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}

	// Players
	playerRepo := &repository.PlayerRepository{DB: database}
	playerService := &services.PlayerService{Repo: playerRepo}
	playerHandler := &handlers.PlayerHandler{Service: playerService}

	// Teams
	teamRepo := &repository.TeamRepository{DB: database}
	teamService := &services.TeamService{Repo: teamRepo}
	teamHandler := &handlers.TeamHandler{Service: teamService}

	mux := http.NewServeMux()

	// registrar swagger
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// rutas normales
	routes.RegisterRoutes(mux, playerHandler, teamHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Servidor corriendo en http://localhost:" + port)

	err = http.ListenAndServe(":"+port, enableCORS(mux))
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
