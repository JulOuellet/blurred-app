package seasons

import "database/sql"

type SeasonRepository interface {
	GetAll() ([]SeasonModel, error)
}

type seasonRepository struct {
	db *sql.DB
}

func NewSeasonRepository(db *sql.DB) SeasonRepository {
	return &seasonRepository{db: db}
}

func (r *seasonRepository) GetAll() ([]SeasonModel, error) {
	rows, err := r.db.Query("SELECT * FROM seasons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seasons []SeasonModel
	for rows.Next() {
		var season SeasonModel
		if err := rows.Scan(
			&season.ID,
			&season.Name,
			&season.StartDate,
			&season.EndDate,
			&season.SportID,
			&season.CreatedAt,
			&season.UpdatedAt,
		); err != nil {
			return nil, err
		}

		seasons = append(seasons, season)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return seasons, nil
}
