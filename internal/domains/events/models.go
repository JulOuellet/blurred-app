package events

import (
	"time"

	"github.com/google/uuid"
)

type EventModel struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	Name           string     `json:"name" db:"name"`
	Date           *time.Time `json:"date" db:"date"`
	ChampionshipID uuid.UUID  `json:"championshipId" db:"championship_id"`
	CreatedAt      time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time  `json:"updatedAt" db:"updated_at"`
}

type EventRequest struct {
	Name           string     `json:"name"`
	Date           *time.Time `json:"date"`
	ChampionshipID uuid.UUID  `json:"championshipId"`
}
