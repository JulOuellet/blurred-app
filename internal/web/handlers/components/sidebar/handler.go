package sidebar

import (
	"net/http"

	"github.com/JulOuellet/blurred-app/internal/domains/sports"
	"github.com/JulOuellet/blurred-app/templates/components/sidebars"
	"github.com/labstack/echo/v4"
)

type SidebarHandler interface {
	GetSidebar(c echo.Context) error
	RefreshSports(c echo.Context) error
}

type sidebarHandler struct {
	sportService sports.SportService
}

func NewSidebarHandler(sportService sports.SportService) SidebarHandler {
	return &sidebarHandler{sportService: sportService}
}

func (h *sidebarHandler) GetSidebar(c echo.Context) error {
	sportsWithSeasons, err := h.sportService.GetAllWithSeasons()
	if err != nil {
		component := sidebars.Sidebar([]sports.SportWithSeasons{})
		return component.Render(c.Request().Context(), c.Response().Writer)
	}

	component := sidebars.Sidebar(sportsWithSeasons)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func (h *sidebarHandler) RefreshSports(c echo.Context) error {
	sportsWithSeasons, err := h.sportService.GetAllWithSeasons()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get sports")
	}

	component := sidebars.SportsList(sportsWithSeasons)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
