package seasons

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SeasonHandler interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
}

type seasonHandler struct {
	seasonService SeasonService
}

func NewSeasonHandler(seasonService SeasonService) SeasonHandler {
	return &seasonHandler{seasonService: seasonService}
}

func (h *seasonHandler) GetAll(c echo.Context) error {
	seasons, err := h.seasonService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get seasons")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"data":  seasons,
		"count": len(seasons),
	})
}

func (h *seasonHandler) GetById(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Season ID is required")
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid season ID format")
	}

	season, err := h.seasonService.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Season not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get season")
	}

	return c.JSON(http.StatusOK, map[string]any{"data": season})
}

func (h *seasonHandler) Create(c echo.Context) error {
	var req SeasonRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	season, err := h.seasonService.Create(req)
	if err != nil {
		c.Logger().Errorf("failed to create season: %v", err)
		if strings.Contains(err.Error(), "cannot be empty") {
			return echo.NewHTTPError(http.StatusBadRequest, "Season name is required")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create season")
	}

	return c.JSON(http.StatusCreated, map[string]any{"data": season})
}
