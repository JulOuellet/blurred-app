package seasons

import "github.com/jmoiron/sqlx"

type SeasonRepository interface {
	GetAll() ([]SeasonModel, error)
}

type seasonRepository struct {
	db *sqlx.DB
}

func NewSeasonRepository(db *sqlx.DB) SeasonRepository {
	return &seasonRepository{db: db}
}

func (r *seasonRepository) GetAll() ([]SeasonModel, error) {
	query := `SELECT id, name, start_date`
}
