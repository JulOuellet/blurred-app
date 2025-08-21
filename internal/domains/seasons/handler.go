package seasons

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SeasonHandler interface {
	GetAll(c echo.Context) error
}

type seasonHandler struct {
	service SeasonSerivce
}

func NewSeasonHandler(service SeasonSerivce) SeasonHandler {
	return &seasonHandler{service: service}
}

func (h *seasonHandler) GetAll(c echo.Context) error {
	seasons, err := h.service.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch seasons")
	}

	return c.JSON(http.StatusOK, seasons)
}
