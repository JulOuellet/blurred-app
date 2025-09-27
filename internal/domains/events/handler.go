package events

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EventHandler interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	GetAllByChampionshipId(c echo.Context) error
}

type eventHandler struct {
	eventService EventService
}

func NewEventHandler(eventService EventService) EventHandler {
	return &eventHandler{eventService: eventService}
}

func (h *eventHandler) GetAll(c echo.Context) error {
	events, err := h.eventService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get events")
	}
	return c.JSON(http.StatusOK, map[string]any{
		"data":  events,
		"count": len(events),
	})
}

func (h *eventHandler) GetById(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Event ID is required")
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid event ID format")
	}

	event, err := h.eventService.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Event not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get event")
	}

	return c.JSON(http.StatusOK, map[string]any{"data": event})
}

func (h *eventHandler) GetAllByChampionshipId(c echo.Context) error {
	championshipIdParam := c.Param("championship_id")
	if championshipIdParam == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Championship ID is required")
	}

	championshipId, err := uuid.Parse(championshipIdParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid championship ID format")
	}

	events, err := h.eventService.GetAllByChampionshipId(championshipId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get events")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"data":  events,
		"count": len(events),
	})
}

func (h *eventHandler) Create(c echo.Context) error {
	var req EventRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	event, err := h.eventService.Create(req)
	if err != nil {
		c.Logger().Errorf("failed to create event: %v", err)
		if strings.Contains(err.Error(), "name cannot be empty") {
			return echo.NewHTTPError(http.StatusBadRequest, "Event name is required")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create event")
	}

	return c.JSON(http.StatusCreated, map[string]any{"data": event})
}
