package pages

import (
	"strings"

	"github.com/JulOuellet/blurred-app/internal/domains/championships"
	"github.com/JulOuellet/blurred-app/internal/domains/events"
	"github.com/JulOuellet/blurred-app/internal/domains/highlights"
	"github.com/JulOuellet/blurred-app/internal/domains/seasons"
)

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
