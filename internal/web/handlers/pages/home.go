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

// recentEventsFetchLimit over-fetches so the "Catch up" section can still be
// filled after events already covered by the hero are filtered out.
const recentEventsFetchLimit = recentEventsLimit * 3

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
	recentEvents, err := h.eventService.GetRecentWithHighlights(recentEventsFetchLimit)
	if err != nil {
		return err
	}

	hero, err := h.buildHero(ongoingChampionships)
	if err != nil {
		return err
	}
	recentEvents = catchUpEvents(recentEvents, hero)

	// Ongoing races beyond the one in the hero still deserve a spot in the
	// "Coming up" strip rather than disappearing entirely.
	if len(ongoingChampionships) > 1 {
		upcomingChampionships = append(ongoingChampionships[1:], upcomingChampionships...)
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		err = pages.HomeContent(hero, upcomingChampionships, recentEvents).Render(c.Request().Context(), c.Response().Writer)
		if err != nil {
			return err
		}
		return sidebars.Sidebar(sportsList, "", true).Render(c.Request().Context(), c.Response().Writer)
	}
	return pages.HomePage(sportsList, hero, upcomingChampionships, recentEvents).Render(c.Request().Context(), c.Response().Writer)
}

// buildHero assembles the "happening now" block from the first ongoing
// championship: its events sorted chronologically (the stage strip) and the
// most recent one that already has highlights (the primary call to action).
func (h *HomePageHandler) buildHero(ongoing []championships.HomeChampionship) (*pages.HomeHero, error) {
	if len(ongoing) == 0 {
		return nil, nil
	}
	champ := ongoing[0]
	champEvents, err := h.eventService.GetAllByChampionshipId(champ.ID, events.SortByDate, events.SortDirectionAsc)
	if err != nil {
		return nil, err
	}
	hero := &pages.HomeHero{Championship: champ, Events: champEvents}
	for i := range champEvents {
		if champEvents[i].HighlightCount > 0 {
			hero.Latest = &champEvents[i]
		}
	}
	return hero, nil
}

// catchUpEvents drops recent events already covered by the hero's stage strip
// and trims the rest down to the display limit.
func catchUpEvents(recent []events.RecentEvent, hero *pages.HomeHero) []events.RecentEvent {
	filtered := make([]events.RecentEvent, 0, len(recent))
	for _, ev := range recent {
		if hero != nil && ev.ChampionshipID == hero.Championship.ID {
			continue
		}
		filtered = append(filtered, ev)
		if len(filtered) == recentEventsLimit {
			break
		}
	}
	return filtered
}
