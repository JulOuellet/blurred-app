package championships

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ChampionshipHandler interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
}

type championshipHandler struct {
	championshipService ChampionshipService
}

func NewChampionshipHandler(championshipService ChampionshipService) ChampionshipHandler {
	return &championshipHandler{championshipService: championshipService}
}

func (h *championshipHandler) GetAll(c echo.Context) error {
	championships, err := h.championshipService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get championships")
	}
	return c.JSON(http.StatusOK, map[string]any{
		"data":  championships,
		"count": len(championships),
	})
}

func (h *championshipHandler) GetById(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Championship ID is required")
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid championship ID format")
	}

	championship, err := h.championshipService.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Championship not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get championship")
	}

	return c.JSON(http.StatusOK, map[string]any{"data": championship})
}

func (h *championshipHandler) Create(c echo.Context) error {
	var req ChampionshipRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	championship, err := h.championshipService.Create(req)
	if err != nil {
		c.Logger().Errorf("failed to create championship: %v", err)

		if strings.Contains(err.Error(), "cannot be empty") {
			if strings.Contains(err.Error(), "name") {
				return echo.NewHTTPError(http.StatusBadRequest, "Championship name is required")
			}
			if strings.Contains(err.Error(), "reference image URL") {
				return echo.NewHTTPError(http.StatusBadRequest, "Reference image URL is required")
			}
		}

		if strings.Contains(err.Error(), "end date cannot be before start date") {
			return echo.NewHTTPError(http.StatusBadRequest, "End date cannot be before start date")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create championship")
	}

	return c.JSON(http.StatusCreated, map[string]any{"data": championship})
}
