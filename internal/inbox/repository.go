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
	List(status string, limit int) ([]InboxItemWithChannel, error)
	CountsByStatus() ([]StatusCount, error)
	Retry(id uuid.UUID) error
	GetById(id uuid.UUID) (*InboxItem, error)
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

func (r *inboxRepository) GetById(id uuid.UUID) (*InboxItem, error) {
	query := `
		SELECT
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
		FROM youtube_inbox
		WHERE id = $1
	`
	var item InboxItem
	if err := r.db.Get(&item, query, id); err != nil {
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

func (r *inboxRepository) List(status string, limit int) ([]InboxItemWithChannel, error) {
	query := `
		SELECT
		  yi.id,
		  yi.integration_id,
		  yi.youtube_video_id,
		  yi.video_title,
		  yi.published_at,
		  yi.status,
		  yi.failure_reason,
		  yi.retry_count,
		  yi.processed_at,
		  yi.created_at,
		  yi.updated_at,
		  i.youtube_channel_name AS channel_name
		FROM youtube_inbox yi
		JOIN integrations i ON i.id = yi.integration_id
		WHERE ($1 = '' OR yi.status = $1)
		ORDER BY yi.created_at DESC
		LIMIT $2
	`
	var items []InboxItemWithChannel
	err := r.db.Select(&items, query, status, limit)
	return items, err
}

func (r *inboxRepository) CountsByStatus() ([]StatusCount, error) {
	query := `
		SELECT status, COUNT(*)::int AS count
		FROM youtube_inbox
		GROUP BY status
	`
	var counts []StatusCount
	err := r.db.Select(&counts, query)
	return counts, err
}

// Retry re-queues a finished item so the processor picks it up again,
// e.g. after fixing the integration pattern that made it skip or fail.
func (r *inboxRepository) Retry(id uuid.UUID) error {
	query := `
		UPDATE youtube_inbox
		SET
		  status = $2,
		  retry_count = 0,
		  failure_reason = NULL,
		  processed_at = NULL
		WHERE id = $1 AND status IN ($3, $4, $5)
	`
	_, err := r.db.Exec(query, id, StatusPending, StatusSkipped, StatusFailed, StatusDead)
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
