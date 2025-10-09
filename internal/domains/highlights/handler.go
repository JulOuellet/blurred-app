package highlights

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type HighlightHandler interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	GetAllByEventId(c echo.Context) error
}

type highlightHandler struct {
	highlightService HighlightService
}

func NewHighlightHandler(highlightService HighlightService) HighlightHandler {
	return &highlightHandler{highlightService: highlightService}
}

func (h *highlightHandler) GetAll(c echo.Context) error {
	highlights, err := h.highlightService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get highlights")
	}
	return c.JSON(http.StatusOK, map[string]any{
		"data":  highlights,
		"count": len(highlights),
	})
}

func (h *highlightHandler) GetById(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Highlight ID is required")
	}
	id, err := uuid.Parse(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid highlight ID format")
	}
	highlight, err := h.highlightService.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Highlight not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get highlight")
	}
	return c.JSON(http.StatusOK, map[string]any{"data": highlight})
}

func (h *highlightHandler) GetAllByEventId(c echo.Context) error {
	eventIdParam := c.Param("event_id")
	if eventIdParam == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Event ID is required")
	}
	eventId, err := uuid.Parse(eventIdParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid event ID format")
	}
	highlights, err := h.highlightService.GetAllByEventId(eventId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get highlights")
	}
	return c.JSON(http.StatusOK, map[string]any{
		"data":  highlights,
		"count": len(highlights),
	})
}

func (h *highlightHandler) Create(c echo.Context) error {
	var req HighlightRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	highlight, err := h.highlightService.Create(req)
	if err != nil {
		c.Logger().Errorf("failed to create highlight: %v", err)
		if strings.Contains(err.Error(), "name cannot be empty") {
			return echo.NewHTTPError(http.StatusBadRequest, "Highlight name is required")
		}
		if strings.Contains(err.Error(), "url cannot be empty") {
			return echo.NewHTTPError(http.StatusBadRequest, "Highlight url is required")
		}
		if strings.Contains(err.Error(), "language cannot be empty") {
			return echo.NewHTTPError(http.StatusBadRequest, "Highlight language is required")
		}
		if strings.Contains(err.Error(), "media type cannot be empty") {
			return echo.NewHTTPError(http.StatusBadRequest, "Highlight media type is required")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create highlight")
	}
	return c.JSON(http.StatusCreated, map[string]any{"data": highlight})
}
