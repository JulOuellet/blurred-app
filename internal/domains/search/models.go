package search

import "github.com/google/uuid"

type SearchResult struct {
	Type  string    `db:"type"`
	ID    uuid.UUID `db:"id"`
	Name  string    `db:"name"`
	URL   string
	Extra string `db:"extra"`
}
