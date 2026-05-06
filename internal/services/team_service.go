package services

import (
	"nbatracker-backend/internal/models"
	"nbatracker-backend/internal/repository"
)

type TeamService struct {
	Repo *repository.TeamRepository
}

func (s *TeamService) GetTeams() ([]models.Team, error) {
	return s.Repo.GetAll()
}
