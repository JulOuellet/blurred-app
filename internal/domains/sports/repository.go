package sports

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SportRepository interface {
	GetAll() ([]SportModel, error)
	GetById(id uuid.UUID) (*SportModel, error)
	Create(name string) (*SportModel, error)
	Update(id uuid.UUID, name string) (*SportModel, error)
}

type sportRepository struct {
	db *sqlx.DB
}

func NewSportRepository(db *sqlx.DB) SportRepository {
	return &sportRepository{db: db}
}

func (r *sportRepository) GetAll() ([]SportModel, error) {
	query := `SELECT id, name, created_at, updated_at 
			  FROM sports 
			  ORDER BY created_at DESC`

	var sports []SportModel
	return sports, r.db.Select(&sports, query)
}

func (r *sportRepository) GetById(id uuid.UUID) (*SportModel, error) {
	query := `SELECT id, name, created_at, updated_at 
			  FROM sports 
			  WHERE id = $1`

	var sport SportModel
	err := r.db.Get(&sport, query, id)
	if err != nil {
		return nil, err
	}

	return &sport, nil
}

func (r *sportRepository) Create(name string) (*SportModel, error) {
	query := `INSERT INTO sports (name)
			  VALUES ($1)
	          RETURNING id, name, created_at, updated_at`

	var sport SportModel
	err := r.db.Get(&sport, query, name)
	if err != nil {
		return nil, err
	}

	return &sport, nil
}

func (r *sportRepository) Update(id uuid.UUID, name string) (*SportModel, error) {
	query := `UPDATE sports
			  SET name = $1
			  WHERE id = $2
			  RETURNING id, name, created_at, updated_at`

	var sport SportModel
	err := r.db.Get(&sport, query, name, id)
	if err != nil {
		return nil, err
	}

	return &sport, nil
}
