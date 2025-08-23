package sports

import (
	"time"

	"github.com/google/uuid"
)

type SportModel struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateSportRequest struct {
	Name string `json:"name"`
}
