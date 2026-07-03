package pages

import (
	"github.com/JulOuellet/blurred-app/internal/domains/championships"
	"github.com/JulOuellet/blurred-app/internal/domains/events"
	"github.com/JulOuellet/blurred-app/internal/domains/sports"
	"github.com/JulOuellet/blurred-app/templates/components/sidebars"
	"github.com/JulOuellet/blurred-app/templates/pages"
	"github.com/labstack/echo/v4"
)

const recentEventsLimit = 6

type HomePageHandler struct {
	sportService        sports.SportService
	championshipService championships.ChampionshipService
	eventService        events.EventService
}

func NewHomePageHandler(
	sportService sports.SportService,
	championshipService championships.ChampionshipService,
	eventService events.EventService,
) HomePageHandler {
	return HomePageHandler{
		sportService:        sportService,
		championshipService: championshipService,
		eventService:        eventService,
	}
}

func (h *HomePageHandler) GetHome(c echo.Context) error {
	sportsList, err := h.sportService.GetAllWithSeasons()
	if err != nil {
		return err
	}
	ongoingChampionships, err := h.championshipService.GetOngoing()
	if err != nil {
		return err
	}
	upcomingChampionships, err := h.championshipService.GetUpcoming(10)
	if err != nil {
		return err
	}
	recentEvents, err := h.eventService.GetRecentWithHighlights(recentEventsLimit)
	if err != nil {
		return err
	}
	if c.Request().Header.Get("HX-Request") == "true" {
		err = pages.HomeContent(sportsList, ongoingChampionships, upcomingChampionships, recentEvents).Render(c.Request().Context(), c.Response().Writer)
		if err != nil {
			return err
		}
		return sidebars.Sidebar(sportsList, "", true).Render(c.Request().Context(), c.Response().Writer)
	}
	return pages.HomePage(sportsList, ongoingChampionships, upcomingChampionships, recentEvents).Render(c.Request().Context(), c.Response().Writer)
}
