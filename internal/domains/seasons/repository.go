package seasons

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SeasonRepository interface {
	GetAll() ([]SeasonModel, error)
	GetById(id uuid.UUID) (*SeasonModel, error)
	Create(name string, startDate time.Time, endDate time.Time, sportId uuid.UUID) (*SeasonModel, error)
}

type seasonRepository struct {
	db *sqlx.DB
}

func NewSeasonRepository(db *sqlx.DB) SeasonRepository {
	return &seasonRepository{db: db}
}

func (r *seasonRepository) GetAll() ([]SeasonModel, error) {
	query := `SELECT id, name, start_date, end_date, sport_id, created_at, updated_at
			  FROM seasons
			  ORDER BY created_at DESC`

	var seasons []SeasonModel
	return seasons, r.db.Select(&seasons, query)
}

func (r *seasonRepository) GetById(id uuid.UUID) (*SeasonModel, error) {
	query := `SELECT id, name, start_date, end_date, sport_id, created_at, updated_at
			  FROM seasons
			  WHERE id = $1`

	var season SeasonModel
	err := r.db.Get(&season, query, id)
	if err != nil {
		return nil, err
	}

	return &season, nil
}

func (r *seasonRepository) Create(
	name string,
	startDate time.Time,
	endDate time.Time,
	sportId uuid.UUID,
) (*SeasonModel, error) {
	query := `INSERT INTO seasons (name, start_date, end_date, sport_id)
			  values ($1, $2, $3, $4)
			  RETURNING id, name, start_date, end_date, sport_id, created_at, updated_at`

	var season SeasonModel
	err := r.db.Get(&season, query, name, startDate, endDate, sportId)
	if err != nil {
		return nil, err
	}
	return &season, nil
}
