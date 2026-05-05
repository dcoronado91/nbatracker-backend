// cmd/server/main.go
package main

import (
    "log"
    "net/http"

    "nbatracker-backend/internal/db"
    "nbatracker-backend/internal/repository"
    "nbatracker-backend/internal/services"
    "nbatracker-backend/internal/handlers"
    "nbatracker-backend/internal/routes"
)

func main() {
    database, err := db.Connect()
    if err != nil {
        log.Fatal(err)
    }

    repo := &repository.PlayerRepository{DB: database}
    service := &services.PlayerService{Repo: repo}
    handler := &handlers.PlayerHandler{Service: service}

    mux := http.NewServeMux()
    routes.RegisterRoutes(mux, handler)

    log.Println("Server running on :8080")
    http.ListenAndServe(":8080", mux)
}