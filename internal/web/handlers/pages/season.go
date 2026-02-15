package pages

import (
	"net/http"

	"github.com/JulOuellet/blurred-app/internal/domains/championships"
	"github.com/JulOuellet/blurred-app/internal/domains/seasons"
	"github.com/JulOuellet/blurred-app/internal/domains/sports"
	"github.com/JulOuellet/blurred-app/templates/components/sidebars"
	"github.com/JulOuellet/blurred-app/templates/pages"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func groupByMonth(champs []championships.ChampionshipModel) []championships.ChampionshipsByMonth {
	var result []championships.ChampionshipsByMonth

	for _, c := range champs {
		month := "TBD"
		if c.StartDate != nil {
			month = c.StartDate.Format("January 2006")
		}

		if len(result) > 0 && result[len(result)-1].Month == month {
			result[len(result)-1].Championships = append(result[len(result)-1].Championships, c)
		} else {
			result = append(result, championships.ChampionshipsByMonth{
				Month:         month,
				Championships: []championships.ChampionshipModel{c},
			})
		}
	}

	return result
}

type SeasonPageHandler struct {
	seasonService   seasons.SeasonService
	championService championships.ChampionshipService
	sportService    sports.SportService
}

func NewSeasonPageHandler(
	seasonService seasons.SeasonService,
	championService championships.ChampionshipService,
	sportService sports.SportService,
) *SeasonPageHandler {
	return &SeasonPageHandler{
		seasonService:   seasonService,
		championService: championService,
		sportService:    sportService,
	}
}

func (h *SeasonPageHandler) GetSeason(c echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid season ID")
	}

	season, err := h.seasonService.GetById(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Season not found")
	}

	champs, err := h.championService.GetAllBySeasonId(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve championships")
	}

	grouped := groupByMonth(champs)

	sport, err := h.sportService.GetById(season.SportID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve sport")
	}

	sportsList, err := h.sportService.GetAllWithSeasons()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve sports list")
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		err = pages.SeasonContent(season, grouped, *sport).Render(c.Request().Context(), c.Response().Writer)
		if err != nil {
			return err
		}
		return sidebars.Sidebar(sportsList, season.ID.String(), true).Render(c.Request().Context(), c.Response().Writer)
	}

	return pages.SeasonPage(season, grouped, *sport, sportsList).Render(c.Request().Context(), c.Response().Writer)
}
