package pages

import (
	"net/http"
	"strings"
	"time"

	"github.com/JulOuellet/blurred-app/internal/domains/championships"
	"github.com/JulOuellet/blurred-app/internal/domains/events"
	"github.com/JulOuellet/blurred-app/internal/domains/seasons"
	"github.com/JulOuellet/blurred-app/templates/layouts"
	"github.com/labstack/echo/v4"
)

type SEOHandler struct {
	seasonService       seasons.SeasonService
	championshipService championships.ChampionshipService
	eventService        events.EventService
}

func NewSEOHandler(
	seasonService seasons.SeasonService,
	championshipService championships.ChampionshipService,
	eventService events.EventService,
) *SEOHandler {
	return &SEOHandler{
		seasonService:       seasonService,
		championshipService: championshipService,
		eventService:        eventService,
	}
}

func (h *SEOHandler) GetRobotsTxt(c echo.Context) error {
	var b strings.Builder
	b.WriteString("User-agent: *\n")
	b.WriteString("Disallow: /admin\n")
	b.WriteString("Disallow: /api\n")
	if base := layouts.PublicBaseURL(); base != "" {
		b.WriteString("Sitemap: " + base + "/sitemap.xml\n")
	}
	return c.String(http.StatusOK, b.String())
}

func (h *SEOHandler) GetSitemap(c echo.Context) error {
	base := layouts.PublicBaseURL()
	if base == "" {
		return echo.NewHTTPError(http.StatusNotFound, "sitemap unavailable")
	}

	allSeasons, err := h.seasonService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load seasons")
	}
	allChampionships, err := h.championshipService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load championships")
	}
	allEvents, err := h.eventService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load events")
	}

	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	b.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n")
	writeURL(&b, base+"/", nil)
	writeURL(&b, base+"/about", nil)
	for _, s := range allSeasons {
		writeURL(&b, base+"/seasons/"+s.ID.String(), &s.UpdatedAt)
	}
	for _, ch := range allChampionships {
		writeURL(&b, base+"/championships/"+ch.ID.String(), &ch.UpdatedAt)
	}
	for _, e := range allEvents {
		writeURL(&b, base+"/events/"+e.ID.String(), &e.UpdatedAt)
	}
	b.WriteString("</urlset>\n")

	return c.Blob(http.StatusOK, "application/xml", []byte(b.String()))
}

func writeURL(b *strings.Builder, loc string, lastMod *time.Time) {
	b.WriteString("  <url><loc>" + loc + "</loc>")
	if lastMod != nil {
		b.WriteString("<lastmod>" + lastMod.Format("2006-01-02") + "</lastmod>")
	}
	b.WriteString("</url>\n")
}
