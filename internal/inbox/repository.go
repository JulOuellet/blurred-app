package inbox

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type InboxRepository interface {
	Insert(integrationID uuid.UUID, videoID string, title string, publishedAt time.Time) error
	ClaimNext() (*InboxItem, error)
	MarkCompleted(id uuid.UUID) error
	MarkSkipped(id uuid.UUID, reason string) error
	MarkFailed(id uuid.UUID, reason string) error
}

type inboxRepository struct {
	db *sqlx.DB
}

func NewInboxRepository(db *sqlx.DB) InboxRepository {
	return &inboxRepository{db: db}
}

func (r *inboxRepository) Insert(integrationID uuid.UUID, videoID string, title string, publishedAt time.Time) error {
	query := `
		INSERT INTO youtube_inbox (
			integration_id,
			youtube_video_id,
			video_title,
			published_at
		)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (youtube_video_id) DO NOTHING
	`
	_, err := r.db.Exec(query, integrationID, videoID, title, publishedAt)
	return err
}

func (r *inboxRepository) ClaimNext() (*InboxItem, error) {
	query := `
		UPDATE youtube_inbox
		SET status = $1
		WHERE id = (
			SELECT id
			FROM youtube_inbox
			WHERE status = $2
			   OR (status = $3 AND retry_count < $4)
			ORDER BY created_at ASC
			FOR UPDATE SKIP LOCKED
			LIMIT 1
		)
		RETURNING
		  id,
		  integration_id,
		  youtube_video_id,
		  video_title,
		  published_at,
		  status,
		  failure_reason,
		  retry_count,
		  processed_at,
		  created_at,
		  updated_at
	`
	var item InboxItem
	err := r.db.Get(&item, query, StatusProcessing, StatusPending, StatusFailed, MaxRetries)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *inboxRepository) MarkCompleted(id uuid.UUID) error {
	query := `
		UPDATE youtube_inbox
		SET status = $2, processed_at = $3
		WHERE id = $1
	`
	_, err := r.db.Exec(query, id, StatusCompleted, time.Now())
	return err
}

func (r *inboxRepository) MarkSkipped(id uuid.UUID, reason string) error {
	query := `
		UPDATE youtube_inbox
		SET status = $2, failure_reason = $3, processed_at = $4
		WHERE id = $1
	`
	_, err := r.db.Exec(query, id, StatusSkipped, reason, time.Now())
	return err
}

func (r *inboxRepository) MarkFailed(id uuid.UUID, reason string) error {
	query := `
		UPDATE youtube_inbox
		SET
		  status = CASE WHEN retry_count + 1 >= $4 THEN $3 ELSE $2 END,
		  failure_reason = $5,
		  retry_count = retry_count + 1
		WHERE id = $1
	`
	_, err := r.db.Exec(query, id, StatusFailed, StatusDead, MaxRetries, reason)
	return err
}
