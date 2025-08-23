package sports

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SportHandler interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
}

type sportHandler struct {
	sportService SportService
}

func NewSportHandler(sportService SportService) SportHandler {
	return &sportHandler{sportService: sportService}
}

func (h *sportHandler) GetAll(c echo.Context) error {
	sports, err := h.sportService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get sports")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"data":  sports,
		"count": len(sports),
	})
}

func (h *sportHandler) GetById(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Sport ID is required")
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid sport ID format")
	}

	sport, err := h.sportService.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Sport not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get sport")
	}

	return c.JSON(http.StatusOK, map[string]any{"data": sport})
}

func (h *sportHandler) Create(c echo.Context) error {
	var req SportRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	sport, err := h.sportService.Create(req)
	if err != nil {
		if isDuplicateKeyError(err) {
			return echo.NewHTTPError(http.StatusConflict, "Sport name already exists")
		}
		if strings.Contains(err.Error(), "cannot be empty") {
			return echo.NewHTTPError(http.StatusBadRequest, "Sport name is required")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create sport")
	}

	return c.JSON(http.StatusCreated, map[string]any{"data": sport})
}

func isDuplicateKeyError(err error) bool {
	errStr := err.Error()
	return strings.Contains(errStr, "duplicate key") ||
		strings.Contains(errStr, "already exists") ||
		strings.Contains(errStr, "23505") // PostgreSQL unique violation
}
