package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type EventRepository interface {
	GetAll() ([]EventModel, error)
	GetById(id uuid.UUID) (*EventModel, error)
	Create(
		name string,
		date *time.Time,
		championshipId uuid.UUID,
	) (*EventModel, error)
	GetAllByChampionshipId(championshipId uuid.UUID) ([]EventModel, error)
}

type eventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) GetAll() ([]EventModel, error) {
	query := `
		SELECT 
		  id, 
		  name, 
		  date, 
		  championship_id, 
		  created_at, 
		  updated_at
		FROM 
		  events
		ORDER BY 
		  created_at DESC
	`
	var events []EventModel
	return events, r.db.Select(&events, query)
}

func (r *eventRepository) GetById(id uuid.UUID) (*EventModel, error) {
	query := `
		SELECT 
		  id, 
		  name, 
		  date, 
		  championship_id, 
		  created_at, 
		  updated_at
		FROM 
		  events
		WHERE
		  id = $1
	`
	var event EventModel
	err := r.db.Get(&event, query, id)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) Create(
	name string,
	date *time.Time,
	championshipId uuid.UUID,
) (*EventModel, error) {
	query := `
		INSERT INTO
		  events (
			name,
			date,
			championship_id
		  )
		VALUES
		  ($1, $2, $3)
		RETURNING
		  id,
		  name,
		  date,
		  championship_id,
		  created_at,
		  updated_at
	`
	var event EventModel
	err := r.db.Get(
		&event,
		query,
		name,
		date,
		championshipId,
	)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) GetAllByChampionshipId(championshipId uuid.UUID) ([]EventModel, error) {
	query := `
	    SELECT
		  id,
		  name,
		  date,
		  championship_id,
		  created_at,
		  updated_at
		FROM
		  events
		WHERE
		  championship_id = $1
		ORDER BY
		  date DESC
	`
	var events []EventModel
	err := r.db.Select(&events, query, championshipId)
	return events, err
}
