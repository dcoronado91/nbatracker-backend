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

func (r *PlayerRepository) Create(p *models.Player) error {
	return r.DB.QueryRow(`
		INSERT INTO players (name, team, image_url, championships, mvp, finals_mvp, dpoy, roty)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`,
		p.Name,
		p.Team,
		p.ImageURL,
		p.Championships,
		p.MVP,
		p.FinalsMVP,
		p.DPOY,
		p.ROTY,
	).Scan(&p.ID)
}

func (r *PlayerRepository) Update(id int, p *models.Player) error {
	_, err := r.DB.Exec(`
		UPDATE players
		SET name=$1, team=$2, image_url=$3,
		    championships=$4, mvp=$5, finals_mvp=$6, dpoy=$7, roty=$8
		WHERE id=$9
	`,
		p.Name,
		p.Team,
		p.ImageURL,
		p.Championships,
		p.MVP,
		p.FinalsMVP,
		p.DPOY,
		p.ROTY,
		id,
	)

	return err
}