package pages

import (
	"net/http"

	"github.com/JulOuellet/blurred-app/internal/domains/championships"
	"github.com/JulOuellet/blurred-app/internal/domains/events"
	"github.com/JulOuellet/blurred-app/internal/domains/seasons"
	"github.com/JulOuellet/blurred-app/internal/domains/sports"
	"github.com/JulOuellet/blurred-app/templates/components/sidebars"
	"github.com/JulOuellet/blurred-app/templates/pages"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ChampionshipPageHandler struct {
	championService championships.ChampionshipService
	seasonService   seasons.SeasonService
	sportService    sports.SportService
	eventService    events.EventService
}

func NewChampionshipPageHandler(
	championService championships.ChampionshipService,
	seasonService seasons.SeasonService,
	sportService sports.SportService,
	eventService events.EventService,
) *ChampionshipPageHandler {
	return &ChampionshipPageHandler{
		championService: championService,
		seasonService:   seasonService,
		sportService:    sportService,
		eventService:    eventService,
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

	evts, err := h.eventService.GetAllByChampionshipId(id, events.SortByDate, events.SortDirectionDesc)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve events")
	}

	sportsList, err := h.sportService.GetAllWithSeasons()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve sports list")
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		err = pages.ChampionshipContent(championship, sport, season, evts).Render(c.Request().Context(), c.Response().Writer)
		if err != nil {
			return err
		}
		return sidebars.Sidebar(sportsList, season.ID.String(), true).Render(c.Request().Context(), c.Response().Writer)
	}

	return pages.ChampionshipPage(championship, sport, season, evts, sportsList).Render(c.Request().Context(), c.Response().Writer)
}

func (h *ChampionshipPageHandler) GetChampionshipEvents(c echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid championship ID")
	}

	sortBy, sortDirection := events.NewSortOptions(
		c.QueryParam("sortBy"),
		c.QueryParam("sortDirection"),
	)

	evts, err := h.eventService.GetAllByChampionshipId(id, sortBy, sortDirection)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve events")
	}

	return pages.EventsList(idStr, evts, sortBy, sortDirection).Render(c.Request().Context(), c.Response().Writer)
}
