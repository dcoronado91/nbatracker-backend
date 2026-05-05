// internal/repository/player_repo.go
package repository

import (
	"database/sql"
	"nbatracker-backend/internal/models"
)

type PlayerRepository struct {
	DB *sql.DB
}

func (r *PlayerRepository) GetAll() ([]models.Player, error) {
	rows, err := r.DB.Query("SELECT * FROM players")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []models.Player

	for rows.Next() {
		var p models.Player
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Team,
			&p.ImageURL,
			&p.Championships,
			&p.MVP,
			&p.FinalsMVP,
			&p.DPOY,
			&p.ROTY,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		players = append(players, p)
	}

	return players, nil
}

func (r *PlayerRepository) GetByID(id int) (*models.Player, error) {
	row := r.DB.QueryRow(`
		SELECT id, name, team, image_url,
		    championships, mvp, finals_mvp, dpoy, roty, created_at
		FROM players
		WHERE id = $1
	`, id)

	var p models.Player

	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Team,
		&p.ImageURL,
		&p.Championships,
		&p.MVP,
		&p.FinalsMVP,
		&p.DPOY,
		&p.ROTY,
		&p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}
