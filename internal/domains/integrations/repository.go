package integrations

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IntegrationRepository interface {
	GetAll() ([]IntegrationModel, error)
	GetAllWithSport() ([]IntegrationWithSport, error)
	GetById(id uuid.UUID) (*IntegrationModel, error)
	GetAllActive() ([]IntegrationModel, error)
	Create(
		youtubeChannelID string,
		youtubeChannelName string,
		sportID uuid.UUID,
		lang string,
		contentFilter *string,
		titleExclude *string,
		stagePattern *string,
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
		  sport_id,
		  lang,
		  content_filter,
		  title_exclude,
		  stage_pattern,
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

func (r *integrationRepository) GetAllWithSport() ([]IntegrationWithSport, error) {
	query := `
		SELECT
		  i.id,
		  i.youtube_channel_id,
		  i.youtube_channel_name,
		  i.sport_id,
		  i.lang,
		  i.content_filter,
		  i.stage_pattern,
		  i.active,
		  i.last_polled_at,
		  i.created_at,
		  i.updated_at,
		  s.name AS sport_name
		FROM
		  integrations i
		JOIN
		  sports s ON s.id = i.sport_id
		ORDER BY
		  s.name ASC, i.created_at DESC
	`
	var integrations []IntegrationWithSport
	return integrations, r.db.Select(&integrations, query)
}

func (r *integrationRepository) GetById(id uuid.UUID) (*IntegrationModel, error) {
	query := `
		SELECT
		  id,
		  youtube_channel_id,
		  youtube_channel_name,
		  sport_id,
		  lang,
		  content_filter,
		  title_exclude,
		  stage_pattern,
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
		  id,
		  youtube_channel_id,
		  youtube_channel_name,
		  sport_id,
		  lang,
		  content_filter,
		  title_exclude,
		  stage_pattern,
		  active,
		  last_polled_at,
		  created_at,
		  updated_at
		FROM
		  integrations
		WHERE
		  active = true
		ORDER BY
		  created_at ASC
	`
	var integrations []IntegrationModel
	return integrations, r.db.Select(&integrations, query)
}

func (r *integrationRepository) Create(
	youtubeChannelID string,
	youtubeChannelName string,
	sportID uuid.UUID,
	lang string,
	contentFilter *string,
	titleExclude *string,
	stagePattern *string,
) (*IntegrationModel, error) {
	query := `
		INSERT INTO
		  integrations (
			youtube_channel_id,
			youtube_channel_name,
			sport_id,
			lang,
			content_filter,
			title_exclude,
			stage_pattern
		  )
		VALUES
		  ($1, $2, $3, $4, $5, $6, $7)
		RETURNING
		  id,
		  youtube_channel_id,
		  youtube_channel_name,
		  sport_id,
		  lang,
		  content_filter,
		  title_exclude,
		  stage_pattern,
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
		sportID,
		lang,
		contentFilter,
		titleExclude,
		stagePattern,
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
