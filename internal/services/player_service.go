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

func (s *PlayerService) GetPlayerByID(id int) (*models.Player, error) {
	return s.Repo.GetByID(id)
}

func (s *PlayerService) CreatePlayer(p *models.Player) error {
	return s.Repo.Create(p)
}