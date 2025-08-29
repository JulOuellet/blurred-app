package seasons

import (
	"time"

	"github.com/google/uuid"
)

type SeasonModel struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	StartDate time.Time `json:"startDate" db:"start_date"`
	EndDate   time.Time `json:"endDate" db:"end_date"`
	SportID   uuid.UUID `json:"sportId" db:"sport_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type SeasonRequest struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	SportID   uuid.UUID `json:"sportId"`
}
