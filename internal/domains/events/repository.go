package events

import (
	"fmt"
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
	GetAllByChampionshipId(championshipId uuid.UUID, sortBy SortBy, sortDirection SortDirection) ([]EventModel, error)
	GetRecentWithHighlights(limit int) ([]RecentEvent, error)
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

func (r *eventRepository) GetAllByChampionshipId(championshipId uuid.UUID, sortBy SortBy, sortDirection SortDirection) ([]EventModel, error) {
	query := fmt.Sprintf(`
	    SELECT
		  e.id,
		  e.name,
		  e.date,
		  e.championship_id,
		  e.created_at,
		  e.updated_at,
		  (SELECT COUNT(*) FROM highlights h WHERE h.event_id = e.id)::int AS highlight_count
		FROM
		  events e
		WHERE
		  e.championship_id = $1
		ORDER BY
		  e.%s %s
	`, sortBy.Column(), sortDirection)
	var events []EventModel
	err := r.db.Select(&events, query, championshipId)
	return events, err
}

// GetRecentWithHighlights returns the events whose highlights arrived most
// recently, for the home page "Latest highlights" section.
func (r *eventRepository) GetRecentWithHighlights(limit int) ([]RecentEvent, error) {
	query := `
		SELECT
		  e.id,
		  e.name,
		  e.date,
		  e.championship_id,
		  e.created_at,
		  e.updated_at,
		  COUNT(h.id)::int AS highlight_count,
		  MAX(h.created_at) AS latest_highlight_at,
		  c.name AS championship_name,
		  sp.name AS sport_name,
		  s.id AS season_id
		FROM events e
		JOIN highlights h ON h.event_id = e.id
		JOIN championships c ON c.id = e.championship_id
		JOIN seasons s ON s.id = c.season_id
		JOIN sports sp ON sp.id = s.sport_id
		GROUP BY e.id, c.name, sp.name, s.id
		ORDER BY latest_highlight_at DESC
		LIMIT $1
	`
	var events []RecentEvent
	err := r.db.Select(&events, query, limit)
	return events, err
}
