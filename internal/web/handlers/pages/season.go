package pages

import (
	"net/http"

	"github.com/JulOuellet/blurred-app/internal/domains/championships"
	"github.com/JulOuellet/blurred-app/internal/domains/seasons"
	"github.com/JulOuellet/blurred-app/internal/domains/sports"
	"github.com/JulOuellet/blurred-app/templates/pages"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SeasonPageHandler struct {
	seasonService   seasons.SeasonService
	championService championships.ChampionshipService
	sportService    sports.SportService
}

func NewSeasonPageHandler(
	seasonService seasons.SeasonService,
	championService championships.ChampionshipService,
	sportService sports.SportService,
) *SeasonPageHandler {
	return &SeasonPageHandler{
		seasonService:   seasonService,
		championService: championService,
		sportService:    sportService,
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

	sport, err := h.sportService.GetById(season.SportID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve sport")
	}

	sportsList, err := h.sportService.GetAllWithSeasons()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve sports list")
	}

	return pages.SeasonPage(season, championships, *sport, sportsList).Render(c.Request().Context(), c.Response().Writer)
}
