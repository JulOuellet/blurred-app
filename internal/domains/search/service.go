package search

import (
	"strings"
)

type SearchService interface {
	Search(query string) ([]SearchResult, error)
}

type searchService struct {
	searchRepo SearchRepository
}

func NewSearchService(searchRepo SearchRepository) SearchService {
	return &searchService{searchRepo: searchRepo}
}

func (s *searchService) Search(query string) ([]SearchResult, error) {
	query = strings.TrimSpace(query)
	if len(query) < 2 {
		return []SearchResult{}, nil
	}

	return s.searchRepo.Search(query, 8)
}
