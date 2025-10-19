package highlights

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type HighlightRepository interface {
	GetAll() ([]HighlightModel, error)
	GetById(id uuid.UUID) (*HighlightModel, error)
	Create(
		name string,
		url string,
		youtubeID string,
		language string,
		mediaType string,
		source string,
		eventID uuid.UUID,
	) (*HighlightModel, error)
	GetAllByEventId(eventId uuid.UUID) ([]HighlightModel, error)
}

type highlightRepository struct {
	db *sqlx.DB
}

func NewHighlightRepository(db *sqlx.DB) HighlightRepository {
	return &highlightRepository{db: db}
}

func (r *highlightRepository) GetAll() ([]HighlightModel, error) {
	query := `
		SELECT 
		  id, 
		  name, 
		  url, 
		  youtube_id,
		  lang, 
		  media_type, 
		  source, 
		  event_id, 
		  created_at, 
		  updated_at
		FROM 
		  highlights
		ORDER BY 
		  created_at DESC
	`
	var highlights []HighlightModel
	return highlights, r.db.Select(&highlights, query)
}

func (r *highlightRepository) GetById(id uuid.UUID) (*HighlightModel, error) {
	query := `
		SELECT 
		  id, 
		  name, 
		  url,
		  youtube_id,
		  lang, 
		  media_type, 
		  source, 
		  event_id, 
		  created_at, 
		  updated_at
		FROM 
		  highlights
		WHERE
		  id = $1
	`
	var highlight HighlightModel
	err := r.db.Get(&highlight, query, id)
	if err != nil {
		return nil, err
	}
	return &highlight, nil
}

func (r *highlightRepository) Create(
	name string,
	url string,
	youtubeID string,
	language string,
	mediaType string,
	source string,
	eventID uuid.UUID,
) (*HighlightModel, error) {
	query := `
		INSERT INTO
		  highlights (
			name,
			url,
		    youtube_id,
			lang,
			media_type,
			source,
			event_id
		  )
		VALUES
		  ($1, $2, $3, $4, $5, $6, $7)
		RETURNING
		  id,
		  name,
		  url,
		  youtube_id,
		  lang,
		  media_type,
		  source,
		  event_id,
		  created_at,
		  updated_at
	`
	var highlight HighlightModel
	err := r.db.Get(
		&highlight,
		query,
		name,
		url,
		youtubeID,
		language,
		mediaType,
		source,
		eventID,
	)
	if err != nil {
		return nil, err
	}
	return &highlight, nil
}

func (r *highlightRepository) GetAllByEventId(eventId uuid.UUID) ([]HighlightModel, error) {
	query := `
	    SELECT
		  id,
		  name,
		  url,
		  youtube_id,
		  lang,
		  media_type,
		  source,
		  event_id,
		  created_at,
		  updated_at
		FROM
		  highlights
		WHERE
		  event_id = $1
		ORDER BY
		  created_at DESC
	`
	var highlights []HighlightModel
	err := r.db.Select(&highlights, query, eventId)
	return highlights, err
}
