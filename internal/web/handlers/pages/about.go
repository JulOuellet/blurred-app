package pages

import (
	"github.com/JulOuellet/blurred-app/internal/domains/sports"
	"github.com/JulOuellet/blurred-app/templates/components/sidebars"
	"github.com/JulOuellet/blurred-app/templates/pages"
	"github.com/labstack/echo/v4"
)

type AboutPageHandler struct {
	sportService sports.SportService
}

func NewAboutPageHandler(sportService sports.SportService) AboutPageHandler {
	return AboutPageHandler{
		sportService: sportService,
	}
}

func (h *AboutPageHandler) GetAbout(c echo.Context) error {
	sportsList, err := h.sportService.GetAllWithSeasons()
	if err != nil {
		return err
	}
	if c.Request().Header.Get("HX-Request") == "true" {
		err = pages.AboutContent().Render(c.Request().Context(), c.Response().Writer)
		if err != nil {
			return err
		}
		return sidebars.Sidebar(sportsList, "", true).Render(c.Request().Context(), c.Response().Writer)
	}
	return pages.AboutPage(sportsList).Render(c.Request().Context(), c.Response().Writer)
}
