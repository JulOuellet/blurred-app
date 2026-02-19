package inbox

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusPending    = "pending"
	StatusProcessing = "processing"
	StatusCompleted  = "completed"
	StatusSkipped    = "skipped"
	StatusFailed     = "failed"
	StatusDead       = "dead"

	MaxRetries = 3
)

type InboxItem struct {
	ID             uuid.UUID  `db:"id"`
	IntegrationID  uuid.UUID  `db:"integration_id"`
	YoutubeVideoID string     `db:"youtube_video_id"`
	VideoTitle     string     `db:"video_title"`
	PublishedAt    *time.Time `db:"published_at"`
	Status         string     `db:"status"`
	FailureReason  *string    `db:"failure_reason"`
	RetryCount     int        `db:"retry_count"`
	ProcessedAt    *time.Time `db:"processed_at"`
	CreatedAt      time.Time  `db:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at"`
}
