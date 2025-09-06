package pages

import (
	"net/http"

	"github.com/JulOuellet/sportlight/internal/domains/championships"
	"github.com/JulOuellet/sportlight/internal/domains/seasons"
	"github.com/JulOuellet/sportlight/templates/pages"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SeasonPageHandler struct {
	seasonService   seasons.SeasonService
	championService championships.ChampionshipService
}

func NewSeasonPageHandler(
	seasonService seasons.SeasonService,
	championService championships.ChampionshipService,
) *SeasonPageHandler {
	return &SeasonPageHandler{
		seasonService:   seasonService,
		championService: championService,
	}
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

	championships, err := h.championService.GetAllBySeasonId(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve championships")
	}

	return pages.SeasonPage(season, championships).Render(c.Request().Context(), c.Response().Writer)
}
