package integrations

import (
	"time"

	"github.com/google/uuid"
)

type IntegrationModel struct {
	ID                 uuid.UUID  `json:"id" db:"id"`
	YoutubeChannelID   string     `json:"youtubeChannelId" db:"youtube_channel_id"`
	YoutubeChannelName *string    `json:"youtubeChannelName" db:"youtube_channel_name"`
	ChampionshipID     uuid.UUID  `json:"championshipId" db:"championship_id"`
	Lang               string     `json:"lang" db:"lang"`
	RelevancePattern   string     `json:"relevancePattern" db:"relevance_pattern"`
	EventPattern       *string    `json:"eventPattern" db:"event_pattern"`
	Active             bool       `json:"active" db:"active"`
	LastPolledAt       *time.Time `json:"lastPolledAt" db:"last_polled_at"`
	CreatedAt          time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt          time.Time  `json:"updatedAt" db:"updated_at"`
}

type IntegrationWithChampionship struct {
	IntegrationModel
	ChampionshipName string `db:"championship_name"`
}

type IntegrationRequest struct {
	YoutubeChannelID   string    `json:"youtubeChannelId"`
	YoutubeChannelName string    `json:"youtubeChannelName"`
	ChampionshipID     uuid.UUID `json:"championshipId"`
	Lang               string    `json:"lang"`
	RelevancePattern   string    `json:"relevancePattern"`
	EventPattern       string    `json:"eventPattern"`
}
