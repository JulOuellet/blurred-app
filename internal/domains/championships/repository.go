package championships

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ChampionshipRepository interface {
	GetAll() ([]ChampionshipModel, error)
	GetById(id uuid.UUID) (*ChampionshipModel, error)
	Create(
		name string,
		organization string,
		startDate *time.Time,
		endDate *time.Time,
		seasonId uuid.UUID,
		description string,
		referenceImgURL string,
	) (*ChampionshipModel, error)
	GetAllBySeasonId(seasonId uuid.UUID) ([]ChampionshipModel, error)
	GetOngoing() ([]HomeChampionship, error)
	GetUpcoming(limit int) ([]HomeChampionship, error)
}

type championshipRepository struct {
	db *sqlx.DB
}

func NewChampionshipRepository(db *sqlx.DB) ChampionshipRepository {
	return &championshipRepository{db: db}
}

func (r *championshipRepository) GetAll() ([]ChampionshipModel, error) {
	query := `
		SELECT 
		  id, 
		  name, 
		  organization, 
		  start_date, 
		  end_date, 
		  season_id, 
		  description,
		  reference_img_url,
		  created_at, 
		  updated_at
		FROM 
		  championships
		ORDER BY 
		  created_at DESC
	`

	var championships []ChampionshipModel
	return championships, r.db.Select(&championships, query)
}

func (r *championshipRepository) GetById(id uuid.UUID) (*ChampionshipModel, error) {
	query := `
		SELECT 
		  id, 
		  name, 
		  organization, 
		  start_date, 
		  end_date, 
		  season_id, 
		  description,
		  reference_img_url,
		  created_at, 
		  updated_at
		FROM 
		  championships
		WHERE
		  id = $1
	`

	var championship ChampionshipModel
	err := r.db.Get(&championship, query, id)
	if err != nil {
		return nil, err
	}

	return &championship, nil
}

func (r *championshipRepository) Create(
	name string,
	organization string,
	startDate *time.Time,
	endDate *time.Time,
	seasonId uuid.UUID,
	description string,
	referenceImgURL string,
) (*ChampionshipModel, error) {
	query := `
		INSERT INTO
		  championships (
			name,
			organization,
			start_date,
			end_date,
			season_id,
			description,
			reference_img_url
		  )
		VALUES
		  ($1, $2, $3, $4, $5, $6, $7)
		RETURNING
		  id,
		  name,
		  organization,
		  start_date,
		  end_date,
		  season_id,
		  description,
		  reference_img_url,
		  created_at,
		  updated_at
	`

	var championship ChampionshipModel
	err := r.db.Get(
		&championship,
		query,
		name,
		organization,
		startDate,
		endDate,
		seasonId,
		description,
		referenceImgURL,
	)
	if err != nil {
		return nil, err
	}

	return &championship, nil
}

func (r *championshipRepository) GetAllBySeasonId(seasonId uuid.UUID) ([]ChampionshipModel, error) {
	query := `
	    SELECT
		  id,
		  name,
		  organization,
		  start_date,
		  end_date,
		  season_id,
		  description,
		  reference_img_url,
		  created_at,
		  updated_at
		FROM
		  championships
		WHERE
		  season_id = $1
		ORDER BY
		  start_date ASC
	`

	var championships []ChampionshipModel
	err := r.db.Select(&championships, query, seasonId)
	return championships, err
}

func (r *championshipRepository) GetOngoing() ([]HomeChampionship, error) {
	query := `
		SELECT
		  c.id,
		  c.name,
		  c.organization,
		  c.start_date,
		  c.end_date,
		  c.season_id,
		  c.description,
		  c.reference_img_url,
		  c.created_at,
		  c.updated_at,
		  sp.name AS sport_name
		FROM championships c
		JOIN seasons s ON c.season_id = s.id
		JOIN sports sp ON s.sport_id = sp.id
		WHERE c.start_date <= NOW() AND c.end_date >= NOW()
		ORDER BY c.end_date ASC
	`

	var championships []HomeChampionship
	return championships, r.db.Select(&championships, query)
}

func (r *championshipRepository) GetUpcoming(limit int) ([]HomeChampionship, error) {
	query := `
		SELECT
		  c.id,
		  c.name,
		  c.organization,
		  c.start_date,
		  c.end_date,
		  c.season_id,
		  c.description,
		  c.reference_img_url,
		  c.created_at,
		  c.updated_at,
		  sp.name AS sport_name
		FROM championships c
		JOIN seasons s ON c.season_id = s.id
		JOIN sports sp ON s.sport_id = sp.id
		WHERE c.start_date > NOW()
		ORDER BY c.start_date ASC
		LIMIT $1
	`

	var championships []HomeChampionship
	return championships, r.db.Select(&championships, query, limit)
}
