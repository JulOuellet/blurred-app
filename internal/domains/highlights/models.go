package highlights

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type HighlightModel struct {
	ID              uuid.UUID `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	GenericName     *string   `json:"genericName" db:"generic_name"`
	URL             string    `json:"url" db:"url"`
	YoutubeID       *string   `json:"youtubeId" db:"youtube_id"`
	DurationSeconds *int      `json:"durationSeconds" db:"duration_seconds"`
	Lang            Language  `json:"lang" db:"lang"`
	MediaType       string    `json:"mediaType" db:"media_type"`
	Source          *string   `json:"source" db:"source"`
	EventID         uuid.UUID `json:"eventId" db:"event_id"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`
}

type HighlightRequest struct {
	Name            string    `json:"name"`
	GenericName     string    `json:"genericName"`
	URL             string    `json:"url"`
	YoutubeID       string    `json:"youtubeId"`
	DurationSeconds int       `json:"durationSeconds"`
	Language        string    `json:"language"`
	MediaType       string    `json:"mediaType"`
	Source          string    `json:"source"`
	EventID         uuid.UUID `json:"eventId"`
}

func (h *HighlightModel) FormatDuration() string {
	if h.DurationSeconds == nil {
		return ""
	}

	seconds := *h.DurationSeconds
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60

	if hours > 0 {
		return fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
	}
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}
