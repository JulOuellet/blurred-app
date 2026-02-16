package championships

import (
	"time"

	"github.com/google/uuid"
)

type ChampionshipModel struct {
	ID              uuid.UUID  `json:"id" db:"id"`
	Name            string     `json:"name" db:"name"`
	Organization    *string    `json:"organization" db:"organization"`
	StartDate       *time.Time `json:"startDate" db:"start_date"`
	EndDate         *time.Time `json:"endDate" db:"end_date"`
	SeasonID        uuid.UUID  `json:"seasonId" db:"season_id"`
	Description     *string    `json:"description" db:"description"`
	ReferenceImgURL string     `json:"referenceImgUrl" db:"reference_img_url"`
	CreatedAt       time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time  `json:"updatedAt" db:"updated_at"`
}

type HomeChampionship struct {
	ChampionshipModel
	SportName string `db:"sport_name"`
}

type ChampionshipsByMonth struct {
	Month         string
	Championships []ChampionshipModel
}

type ChampionshipRequest struct {
	Name            string     `json:"name"`
	Organization    *string    `json:"organization"`
	StartDate       *time.Time `json:"startDate"`
	EndDate         *time.Time `json:"endDate"`
	SeasonID        uuid.UUID  `json:"seasonId"`
	Description     *string    `json:"description"`
	ReferenceImgURL string     `json:"referenceImgUrl"`
}
