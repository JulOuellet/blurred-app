package pages

import (
	"net/http"

	"github.com/JulOuellet/sportlight/internal/domains/championships"
	"github.com/JulOuellet/sportlight/internal/domains/seasons"
	"github.com/JulOuellet/sportlight/internal/domains/sports"
	"github.com/JulOuellet/sportlight/templates/pages"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ChampionshipPageHandler struct {
	championService championships.ChampionshipService
	seasonService   seasons.SeasonService
	sportService    sports.SportService
}

func NewChampionshipPageHandler(
	championService championships.ChampionshipService,
	seasonService seasons.SeasonService,
	sportService sports.SportService,
) *ChampionshipPageHandler {
	return &ChampionshipPageHandler{
		championService: championService,
		seasonService:   seasonService,
		sportService:    sportService,
	}
}

func (h *ChampionshipPageHandler) GetChampionship(c echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid championship ID")
	}

	championship, err := h.championService.GetById(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Championship not found")
	}

	season, err := h.seasonService.GetById(championship.SeasonID)
	if err != nil {
		return c.String(http.StatusNotFound, "Season not found")
	}

	sport, err := h.sportService.GetById(season.SportID)
	if err != nil {
		return c.String(http.StatusNotFound, "Sport not found")
	}

	return pages.ChampionshipPage(championship, sport, season).Render(c.Request().Context(), c.Response().Writer)
}
