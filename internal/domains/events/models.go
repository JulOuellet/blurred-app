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

type SortBy string

const (
	SortByDate SortBy = "date"
)

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

var allowedSortByColumns = map[SortBy]string{
	SortByDate: "date",
}

func NewSortOptions(sortBy, sortDirection string) (SortBy, SortDirection) {
	sb := SortBy(sortBy)
	if _, ok := allowedSortByColumns[sb]; !ok {
		sb = SortByDate
	}

	sd := SortDirection(sortDirection)
	if sd != SortDirectionAsc && sd != SortDirectionDesc {
		sd = SortDirectionDesc
	}

	return sb, sd
}

func (sb SortBy) Column() string {
	if col, ok := allowedSortByColumns[sb]; ok {
		return col
	}
	return "date"
}

type RecentEvent struct {
	EventModel
	ChampionshipName string    `db:"championship_name"`
	SportName        string    `db:"sport_name"`
	SeasonID         uuid.UUID `db:"season_id"`
}
