package pages

import (
	"github.com/JulOuellet/blurred-app/internal/domains/sports"
	"github.com/JulOuellet/blurred-app/templates/pages"
	"github.com/labstack/echo/v4"
)

type HomePageHandler struct {
	sportService sports.SportService
}

func NewHomePageHandler(sportService sports.SportService) HomePageHandler {
	return HomePageHandler{
		sportService: sportService,
	}
}

func (h *HomePageHandler) GetHome(c echo.Context) error {
	sportsList, err := h.sportService.GetAllWithSeasons()
	if err != nil {
		// TODO: Better error handling
		return err
	}
	if c.Request().Header.Get("HX-Request") == "true" {
		return pages.HomeContent().Render(c.Request().Context(), c.Response().Writer)
	}
	return pages.HomePage(sportsList).Render(c.Request().Context(), c.Response().Writer)
}
