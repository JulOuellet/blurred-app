package sports

import (
	"database/sql"
	"log"
)

type SportRepository interface {
	GetAll() ([]SportModel, error)
}

type sportRepository struct {
	db *sql.DB
}

func NewSportRepository(db *sql.DB) SportRepository {
	return &sportRepository{db: db}
}

func (r *sportRepository) GetAll() ([]SportModel, error) {
	rows, err := r.db.Query("SELECT * FROM sports")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sports []SportModel
	for rows.Next() {
		var sport SportModel
		if err := rows.Scan(&sport.ID, &sport.Name, &sport.CreatedAt); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		sports = append(sports, sport)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sports, nil
}
