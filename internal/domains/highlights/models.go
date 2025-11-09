package highlights

import (
	"time"

	"github.com/google/uuid"
)

type HighlightModel struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	URL       string    `json:"url" db:"url"`
	YoutubeID *string   `json:"youtubeId" db:"youtube_id"`
	Lang      Language  `json:"lang" db:"lang"`
	MediaType string    `json:"mediaType" db:"media_type"`
	Source    *string   `json:"source" db:"source"`
	EventID   uuid.UUID `json:"eventId" db:"event_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type HighlightRequest struct {
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	YoutubeID string    `json:"youtubeId"`
	Language  string    `json:"language"`
	MediaType string    `json:"mediaType"`
	Source    string    `json:"source"`
	EventID   uuid.UUID `json:"eventId"`
}
