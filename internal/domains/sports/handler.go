package sports

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SportHandler interface {
	GetAll(c echo.Context) error
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

	return c.JSON(http.StatusOK, sports)
}
