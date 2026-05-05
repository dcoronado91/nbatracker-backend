package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"nbatracker-backend/internal/db"
	"nbatracker-backend/internal/handlers"
	"nbatracker-backend/internal/repository"
	"nbatracker-backend/internal/routes"
	"nbatracker-backend/internal/services"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar .env, usando variables del sistema")
	}

	// Conectar a la base de datos
	database, err := db.Connect()
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}

	// Inicializar capas
	repo := &repository.PlayerRepository{DB: database}
	service := &services.PlayerService{Repo: repo}
	handler := &handlers.PlayerHandler{Service: service}

	// Configurar rutas
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux, handler)

	// Obtener puerto desde .env o usar 8080 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Servidor corriendo en http://localhost:" + port)

	// Levantar servidor
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}