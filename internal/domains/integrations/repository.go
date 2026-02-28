package integrations

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IntegrationRepository interface {
	GetAll() ([]IntegrationModel, error)
	GetAllWithChampionship() ([]IntegrationWithChampionship, error)
	GetById(id uuid.UUID) (*IntegrationModel, error)
	GetAllActive() ([]IntegrationModel, error)
	Create(
		youtubeChannelID string,
		youtubeChannelName string,
		championshipID uuid.UUID,
		lang string,
		relevancePattern string,
		eventPattern *string,
	) (*IntegrationModel, error)
	Delete(id uuid.UUID) error
	UpdateLastPolledAt(id uuid.UUID, t time.Time) error
}

type integrationRepository struct {
	db *sqlx.DB
}

func NewIntegrationRepository(db *sqlx.DB) IntegrationRepository {
	return &integrationRepository{db: db}
}

func (r *integrationRepository) GetAll() ([]IntegrationModel, error) {
	query := `
		SELECT
		  id,
		  youtube_channel_id,
		  youtube_channel_name,
		  championship_id,
		  lang,
		  relevance_pattern,
		  event_pattern,
		  active,
		  last_polled_at,
		  created_at,
		  updated_at
		FROM
		  integrations
		ORDER BY
		  created_at DESC
	`
	var integrations []IntegrationModel
	return integrations, r.db.Select(&integrations, query)
}

func (r *integrationRepository) GetAllWithChampionship() ([]IntegrationWithChampionship, error) {
	query := `
		SELECT
		  i.id,
		  i.youtube_channel_id,
		  i.youtube_channel_name,
		  i.championship_id,
		  i.lang,
		  i.relevance_pattern,
		  i.event_pattern,
		  i.active,
		  i.last_polled_at,
		  i.created_at,
		  i.updated_at,
		  c.name AS championship_name
		FROM
		  integrations i
		JOIN
		  championships c ON c.id = i.championship_id
		ORDER BY
		  c.name ASC, i.created_at DESC
	`
	var integrations []IntegrationWithChampionship
	return integrations, r.db.Select(&integrations, query)
}

func (r *integrationRepository) GetById(id uuid.UUID) (*IntegrationModel, error) {
	query := `
		SELECT
		  id,
		  youtube_channel_id,
		  youtube_channel_name,
		  championship_id,
		  lang,
		  relevance_pattern,
		  event_pattern,
		  active,
		  last_polled_at,
		  created_at,
		  updated_at
		FROM
		  integrations
		WHERE
		  id = $1
	`
	var integration IntegrationModel
	err := r.db.Get(&integration, query, id)
	if err != nil {
		return nil, err
	}
	return &integration, nil
}

func (r *integrationRepository) GetAllActive() ([]IntegrationModel, error) {
	query := `
		SELECT
		  i.id,
		  i.youtube_channel_id,
		  i.youtube_channel_name,
		  i.championship_id,
		  i.lang,
		  i.relevance_pattern,
		  i.event_pattern,
		  i.active,
		  i.last_polled_at,
		  i.created_at,
		  i.updated_at
		FROM
		  integrations i
		JOIN
		  championships c ON c.id = i.championship_id
		WHERE
		  i.active = true
		  AND (c.end_date IS NULL OR c.end_date > NOW())
		ORDER BY
		  i.created_at ASC
	`
	var integrations []IntegrationModel
	return integrations, r.db.Select(&integrations, query)
}

func (r *integrationRepository) Create(
	youtubeChannelID string,
	youtubeChannelName string,
	championshipID uuid.UUID,
	lang string,
	relevancePattern string,
	eventPattern *string,
) (*IntegrationModel, error) {
	query := `
		INSERT INTO
		  integrations (
			youtube_channel_id,
			youtube_channel_name,
			championship_id,
			lang,
			relevance_pattern,
			event_pattern
		  )
		VALUES
		  ($1, $2, $3, $4, $5, $6)
		RETURNING
		  id,
		  youtube_channel_id,
		  youtube_channel_name,
		  championship_id,
		  lang,
		  relevance_pattern,
		  event_pattern,
		  active,
		  last_polled_at,
		  created_at,
		  updated_at
	`
	var integration IntegrationModel
	err := r.db.Get(
		&integration,
		query,
		youtubeChannelID,
		youtubeChannelName,
		championshipID,
		lang,
		relevancePattern,
		eventPattern,
	)
	if err != nil {
		return nil, err
	}
	return &integration, nil
}

func (r *integrationRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM integrations WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *integrationRepository) UpdateLastPolledAt(id uuid.UUID, t time.Time) error {
	query := `
		UPDATE integrations
		SET last_polled_at = $2
		WHERE id = $1
	`
	_, err := r.db.Exec(query, id, t)
	return err
}
