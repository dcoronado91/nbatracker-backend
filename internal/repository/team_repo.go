package repository

import (
	"database/sql"
	"nbatracker-backend/internal/models"
)

type TeamRepository struct {
	DB *sql.DB
}

func (r *TeamRepository) GetAll() ([]models.Team, error) {
	rows, err := r.DB.Query(`
		SELECT name, city, abbreviation, championships, logo_url, conference, division, created_at
		FROM teams
		ORDER BY name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var t models.Team
		if err := rows.Scan(
			&t.Name,
			&t.City,
			&t.Abbreviation,
			&t.Championships,
			&t.LogoURL,
			&t.Conference,
			&t.Division,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}
	return teams, nil
}
