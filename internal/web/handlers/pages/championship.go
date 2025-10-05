package pages

import (
	"net/http"

	"github.com/JulOuellet/blurred-app/internal/domains/championships"
	"github.com/JulOuellet/blurred-app/internal/domains/events"
	"github.com/JulOuellet/blurred-app/internal/domains/seasons"
	"github.com/JulOuellet/blurred-app/internal/domains/sports"
	"github.com/JulOuellet/blurred-app/templates/pages"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ChampionshipPageHandler struct {
	championService championships.ChampionshipService
	seasonService   seasons.SeasonService
	sportService    sports.SportService
	evetService     events.EventService
}

func NewChampionshipPageHandler(
	championService championships.ChampionshipService,
	seasonService seasons.SeasonService,
	sportService sports.SportService,
	evetService events.EventService,
) *ChampionshipPageHandler {
	return &ChampionshipPageHandler{
		championService: championService,
		seasonService:   seasonService,
		sportService:    sportService,
		evetService:     evetService,
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

	events, err := h.evetService.GetAllByChampionshipId(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve events")
	}

	return pages.ChampionshipPage(championship, sport, season, events).Render(c.Request().Context(), c.Response().Writer)
}
