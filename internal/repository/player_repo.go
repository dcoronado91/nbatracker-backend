// internal/repository/player_repo.go
package repository

import (
	"database/sql"
	"nbatracker-backend/internal/models"
	"strconv"
)

type PlayerRepository struct {
	DB *sql.DB
}

func (r *PlayerRepository) GetAll() ([]models.Player, error) {
	rows, err := r.DB.Query("SELECT * FROM players ORDER BY id")
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

func (r *PlayerRepository) GetPaginated(page, limit int) ([]models.Player, int, error) {
	var total int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM players").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	rows, err := r.DB.Query(`
		SELECT id, name, team, image_url,
		    championships, mvp, finals_mvp, dpoy, roty, created_at
		FROM players
		ORDER BY id
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, 0, err
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
			return nil, 0, err
		}
		players = append(players, p)
	}

	return players, total, nil
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

func (r *PlayerRepository) Delete(id int) error {
	result, err := r.DB.Exec(`
		DELETE FROM players
		WHERE id = $1
	`, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *PlayerRepository) GetAdvanced(page, limit int, q, sort, order string) ([]models.Player, int, error) {

	validSort := map[string]string{
		"name":          "name",
		"mvp":           "mvp",
		"championships": "championships",
	}

	sortField, ok := validSort[sort]
	if !ok {
		sortField = "id"
	}

	if order != "desc" {
		order = "asc"
	}

	where := ""
	args := []interface{}{}
	argIndex := 1

	if q != "" {
		where = "WHERE name ILIKE $" + strconv.Itoa(argIndex)
		args = append(args, "%"+q+"%")
		argIndex++
	}

	var total int
	countQuery := "SELECT COUNT(*) FROM players " + where
	err := r.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	query := `
		SELECT id, name, team, image_url,
		    championships, mvp, finals_mvp, dpoy, roty, created_at
		FROM players
	` + where + `
		ORDER BY ` + sortField + ` ` + order + `
		LIMIT $` + strconv.Itoa(argIndex) + `
		OFFSET $` + strconv.Itoa(argIndex+1)

	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
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
			return nil, 0, err
		}
		players = append(players, p)
	}

	return players, total, nil
}
