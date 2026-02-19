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
		genericName string,
		url string,
		youtubeID string,
		durationSeconds int,
		language string,
		mediaType string,
		source string,
		eventID uuid.UUID,
	) (*HighlightModel, error)
	GetAllByEventId(eventId uuid.UUID) ([]HighlightModel, error)
	ExistsByYoutubeID(youtubeID string) (bool, error)
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
		  generic_name,
		  url, 
		  youtube_id,
		  duration_seconds,
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
	      generic_name,
		  url,
		  youtube_id,
		  duration_seconds,
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
	genericName string,
	url string,
	youtubeID string,
	durationSeconds int,
	language string,
	mediaType string,
	source string,
	eventID uuid.UUID,
) (*HighlightModel, error) {
	query := `
		INSERT INTO
		  highlights (
			name,
			generic_name,
			url,
		    youtube_id,
			duration_seconds,
			lang,
			media_type,
			source,
			event_id
		  )
		VALUES
		  ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING
		  id,
		  name,
		  generic_name,
		  url,
		  youtube_id,
		  duration_seconds,
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
		genericName,
		url,
		youtubeID,
		durationSeconds,
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
		  generic_name,
		  url,
		  youtube_id,
		  duration_seconds,
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

func (r *highlightRepository) ExistsByYoutubeID(youtubeID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM highlights WHERE youtube_id = $1)`
	var exists bool
	err := r.db.Get(&exists, query, youtubeID)
	return exists, err
}
