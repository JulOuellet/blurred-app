package pages

import (
	"net/http"

	"github.com/JulOuellet/sportlight/internal/domains/seasons"
	"github.com/JulOuellet/sportlight/templates/pages"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SeasonPageHandler struct {
	seasonService seasons.SeasonService
}

func NewSeasonPageHandler(seasonService seasons.SeasonService) *SeasonPageHandler {
	return &SeasonPageHandler{seasonService: seasonService}
}

func (h *SeasonPageHandler) GetSeason(c echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid season ID")
	}

	season, err := h.seasonService.GetById(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Season not found")
	}

	return pages.SeasonPage(season).Render(c.Request().Context(), c.Response().Writer)
}
