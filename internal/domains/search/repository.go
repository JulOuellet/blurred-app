package search

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SearchRepository interface {
	Search(query string, limit int) ([]SearchResult, error)
}

type searchRepository struct {
	db *sqlx.DB
}

func NewSearchRepository(db *sqlx.DB) SearchRepository {
	return &searchRepository{db: db}
}

func (r *searchRepository) Search(query string, limit int) ([]SearchResult, error) {
	sqlQuery := `
		(SELECT 'sport' AS type, id, name, '' AS extra FROM sports WHERE name ILIKE $1 LIMIT 3)
		UNION ALL
		(SELECT 'championship' AS type, id, name, COALESCE(organization, '') AS extra FROM championships WHERE name ILIKE $1 LIMIT 3)
		UNION ALL
		(SELECT 'event' AS type, id, name, '' AS extra FROM events WHERE name ILIKE $1 LIMIT 3)
		ORDER BY type, name
		LIMIT $2
	`

	pattern := fmt.Sprintf("%%%s%%", query)
	var results []SearchResult
	err := r.db.Select(&results, sqlQuery, pattern, limit)
	if err != nil {
		return nil, err
	}

	for i := range results {
		switch results[i].Type {
		case "sport":
			results[i].URL = fmt.Sprintf("/sports/%s", results[i].ID)
		case "championship":
			results[i].URL = fmt.Sprintf("/championships/%s", results[i].ID)
		case "event":
			results[i].URL = fmt.Sprintf("/events/%s", results[i].ID)
		}
	}

	return results, nil
}
