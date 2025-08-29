package sports

import (
	"time"

	"github.com/JulOuellet/sportlight/internal/domains/seasons"
	"github.com/google/uuid"
)

type SportModel struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type SportRequest struct {
	Name string `json:"name"`
}

type SportWithSeasons struct {
	SportModel
	Seasons []seasons.SeasonModel `json:"seasons"`
}
