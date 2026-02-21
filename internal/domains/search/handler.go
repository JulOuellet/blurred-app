package search

import (
	"net/http"

	"github.com/JulOuellet/blurred-app/templates/components/searchbars"
	"github.com/labstack/echo/v4"
)

type SearchHandler struct {
	searchService SearchService
}

func NewSearchHandler(searchService SearchService) *SearchHandler {
	return &SearchHandler{searchService: searchService}
}

func (h *SearchHandler) Search(c echo.Context) error {
	query := c.QueryParam("q")

	results, err := h.searchService.Search(query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to search")
	}

	items := make([]searchbars.SearchResultItem, len(results))
	for i, r := range results {
		items[i] = searchbars.SearchResultItem{
			Type:  r.Type,
			Name:  r.Name,
			URL:   r.URL,
			Extra: r.Extra,
		}
	}

	return searchbars.SearchResults(items).Render(c.Request().Context(), c.Response().Writer)
}
