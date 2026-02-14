package pages

import (
	"github.com/JulOuellet/blurred-app/internal/domains/events"
	"github.com/JulOuellet/blurred-app/internal/domains/sports"
	"github.com/JulOuellet/blurred-app/templates/components/sidebars"
	"github.com/JulOuellet/blurred-app/templates/pages"
	"github.com/labstack/echo/v4"
)

type HomePageHandler struct {
	sportService sports.SportService
	eventService events.EventService
}

func NewHomePageHandler(sportService sports.SportService, eventService events.EventService) HomePageHandler {
	return HomePageHandler{
		sportService: sportService,
		eventService: eventService,
	}
}

func (h *HomePageHandler) GetHome(c echo.Context) error {
	sportsList, err := h.sportService.GetAllWithSeasons()
	if err != nil {
		return err
	}
	recentEvents, err := h.eventService.GetRecent(10)
	if err != nil {
		return err
	}
	if c.Request().Header.Get("HX-Request") == "true" {
		err = pages.HomeContent(sportsList, recentEvents).Render(c.Request().Context(), c.Response().Writer)
		if err != nil {
			return err
		}
		return sidebars.Sidebar(sportsList, "", true).Render(c.Request().Context(), c.Response().Writer)
	}
	return pages.HomePage(sportsList, recentEvents).Render(c.Request().Context(), c.Response().Writer)
}
