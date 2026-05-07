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

func (s *PlayerService) GetPlayersPaginated(page, limit int) ([]models.Player, int, error) {
    return s.Repo.GetPaginated(page, limit)
}

func (s *PlayerService) GetPlayerByID(id int) (*models.Player, error) {
	return s.Repo.GetByID(id)
}

func (s *PlayerService) CreatePlayer(p *models.Player) error {
	return s.Repo.Create(p)
}

func (s *PlayerService) UpdatePlayer(id int, p *models.Player) error {
	return s.Repo.Update(id, p)
}

func (s *PlayerService) DeletePlayer(id int) error {
	return s.Repo.Delete(id)
}

func (s *PlayerService) GetPlayersAdvanced(page, limit int, q, sort, order string) ([]models.Player, int, error) {
	return s.Repo.GetAdvanced(page, limit, q, sort, order)
}