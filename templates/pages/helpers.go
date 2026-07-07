package pages

import (
	"strings"
	"time"

	"github.com/JulOuellet/blurred-app/internal/domains/championships"
	"github.com/JulOuellet/blurred-app/internal/domains/events"
	"github.com/JulOuellet/blurred-app/internal/domains/highlights"
	"github.com/JulOuellet/blurred-app/internal/domains/seasons"
)

// HomeHero is the "happening now" block at the top of the home page: the
// ongoing championship, its full list of events (sorted by date ascending,
// forming the stage strip) and the most recent event that already has
// highlights, which becomes the primary call to action.
type HomeHero struct {
	Championship championships.HomeChampionship
	Events       []events.EventModel
	Latest       *events.EventModel
}

// relativeDate renders an event date the way a fan reads it: "Today",
// "Yesterday", "Tomorrow", then "June 21" within the current year and a full
// date beyond that.
func relativeDate(t *time.Time) string {
	if t == nil {
		return ""
	}
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	day := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, now.Location())
	switch int(today.Sub(day).Hours() / 24) {
	case 0:
		return "Today"
	case 1:
		return "Yesterday"
	case -1:
		return "Tomorrow"
	}
	if t.Year() == now.Year() {
		return t.Format("January 2")
	}
	return t.Format("January 2, 2006")
}

// isOngoing reports whether a championship is currently running, so the
// "Coming up" strip can mark live races that didn't make the hero.
func isOngoing(c championships.HomeChampionship) bool {
	now := time.Now()
	if c.StartDate == nil || c.StartDate.After(now) {
		return false
	}
	end := c.EndDate
	if end == nil {
		end = c.StartDate
	}
	return end.Year() > now.Year() ||
		(end.Year() == now.Year() && end.YearDay() >= now.YearDay())
}

// heroDateRange formats a championship's running dates for the hero label,
// e.g. "July 4 – 26" or "July 28 – August 3".
func heroDateRange(start, end *time.Time) string {
	if start == nil {
		return ""
	}
	if end == nil || (start.Year() == end.Year() && start.YearDay() == end.YearDay()) {
		return start.Format("January 2")
	}
	if start.Month() == end.Month() {
		return start.Format("January 2") + " – " + end.Format("2")
	}
	return start.Format("January 2") + " – " + end.Format("January 2")
}

// eventTitle builds a search-friendly page title like
// "Tour de France 2026 – Stage 5". The season year is skipped when the
// championship name already contains it.
func eventTitle(championship *championships.ChampionshipModel, season *seasons.SeasonModel, event *events.EventModel) string {
	name := championship.Name
	if !strings.Contains(name, season.Name) {
		name += " " + season.Name
	}
	return name + " – " + event.Name
}

// videoHighlights filters down to the highlights the event page can actually
// render, so the template can show an empty state when nothing is playable.
func videoHighlights(hls []highlights.HighlightModel) []highlights.HighlightModel {
	out := make([]highlights.HighlightModel, 0, len(hls))
	for _, h := range hls {
		if h.MediaType == "VIDEO" && h.YoutubeID != nil {
			out = append(out, h)
		}
	}
	return out
}
