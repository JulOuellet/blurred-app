package sports

import "github.com/google/uuid"

type SportModel struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
