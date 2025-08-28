// internal/web/handlers/sidebar/handler.go
package sidebar

import (
	"net/http"

	"github.com/JulOuellet/sportlight/internal/domains/sports"
	"github.com/JulOuellet/sportlight/templates/components/sidebar"
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
	sportsList, err := h.sportService.GetAll()
	if err != nil {
		component := sidebar.Sidebar([]sports.SportModel{})
		return component.Render(c.Request().Context(), c.Response().Writer)
	}

	component := sidebar.Sidebar(sportsList)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func (h *sidebarHandler) RefreshSports(c echo.Context) error {
	sports, err := h.sportService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get sports")
	}

	component := sidebar.SportsList(sports)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
