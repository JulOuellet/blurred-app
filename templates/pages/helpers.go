package pages

import "github.com/JulOuellet/blurred-app/internal/domains/highlights"

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
