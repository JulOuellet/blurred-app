package integrations

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type IntegrationHandler interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
}

type integrationHandler struct {
	integrationService IntegrationService
}

func NewIntegrationHandler(integrationService IntegrationService) IntegrationHandler {
	return &integrationHandler{integrationService: integrationService}
}

func (h *integrationHandler) GetAll(c echo.Context) error {
	integrations, err := h.integrationService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get integrations")
	}
	return c.JSON(http.StatusOK, map[string]any{
		"data":  integrations,
		"count": len(integrations),
	})
}

func (h *integrationHandler) GetById(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Integration ID is required")
	}
	id, err := uuid.Parse(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid integration ID format")
	}
	integration, err := h.integrationService.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Integration not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get integration")
	}
	return c.JSON(http.StatusOK, map[string]any{"data": integration})
}

func (h *integrationHandler) Create(c echo.Context) error {
	var req IntegrationRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	integration, err := h.integrationService.Create(req)
	if err != nil {
		c.Logger().Errorf("failed to create integration: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]any{"data": integration})
}
