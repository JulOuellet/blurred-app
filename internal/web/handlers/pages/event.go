package pages

import (
	"net/http"

	"github.com/JulOuellet/blurred-app/internal/domains/championships"
	"github.com/JulOuellet/blurred-app/internal/domains/events"
	"github.com/JulOuellet/blurred-app/internal/domains/highlights"
	"github.com/JulOuellet/blurred-app/internal/domains/seasons"
	"github.com/JulOuellet/blurred-app/internal/domains/sports"
	"github.com/JulOuellet/blurred-app/templates/pages"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EventPageHandler struct {
	championService  championships.ChampionshipService
	seasonService    seasons.SeasonService
	sportService     sports.SportService
	eventService     events.EventService
	highlightService highlights.HighlightService
}

func NewEventPageHandler(
	championService championships.ChampionshipService,
	seasonService seasons.SeasonService,
	sportService sports.SportService,
	eventService events.EventService,
	highlightService highlights.HighlightService,
) *EventPageHandler {
	return &EventPageHandler{
		championService:  championService,
		seasonService:    seasonService,
		sportService:     sportService,
		eventService:     eventService,
		highlightService: highlightService,
	}
}

func (h *EventPageHandler) GetEvent(c echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid event ID")
	}

	event, err := h.eventService.GetById(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Event not found")
	}

	championship, err := h.championService.GetById(event.ChampionshipID)
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

	highlights, err := h.highlightService.GetAllByEventId(event.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch highlights")
	}

	sportsList, err := h.sportService.GetAllWithSeasons()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve sports list")
	}

	return pages.EventPage(
		championship,
		sport,
		season,
		event,
		highlights,
		sportsList,
	).Render(c.Request().Context(), c.Response().Writer)
}
