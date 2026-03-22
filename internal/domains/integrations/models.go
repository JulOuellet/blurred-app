package integrations

import (
	"time"

	"github.com/google/uuid"
)

type IntegrationModel struct {
	ID                 uuid.UUID  `json:"id" db:"id"`
	YoutubeChannelID   string     `json:"youtubeChannelId" db:"youtube_channel_id"`
	YoutubeChannelName *string    `json:"youtubeChannelName" db:"youtube_channel_name"`
	SportID            uuid.UUID  `json:"sportId" db:"sport_id"`
	Lang               string     `json:"lang" db:"lang"`
	ContentFilter      *string    `json:"contentFilter" db:"content_filter"`
	TitleExclude       *string    `json:"titleExclude" db:"title_exclude"`
	StagePattern       *string    `json:"stagePattern" db:"stage_pattern"`
	Active             bool       `json:"active" db:"active"`
	LastPolledAt       *time.Time `json:"lastPolledAt" db:"last_polled_at"`
	CreatedAt          time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt          time.Time  `json:"updatedAt" db:"updated_at"`
}

type IntegrationWithSport struct {
	IntegrationModel
	SportName string `db:"sport_name"`
}

type IntegrationRequest struct {
	YoutubeChannelID   string    `json:"youtubeChannelId"`
	YoutubeChannelName string    `json:"youtubeChannelName"`
	SportID            uuid.UUID `json:"sportId"`
	Lang               string    `json:"lang"`
	ContentFilter      string    `json:"contentFilter"`
	TitleExclude       string    `json:"titleExclude"`
	StagePattern       string    `json:"stagePattern"`
}
