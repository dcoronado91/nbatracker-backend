// internal/services/player_service.go
package services

import (
    "nbatracker-backend/internal/models"
    "nbatracker-backend/internal/repository"
)

type PlayerService struct {
    Repo *repository.PlayerRepository
}

func (s *PlayerService) GetPlayers() ([]models.Player, error) {
    return s.Repo.GetAll()
}