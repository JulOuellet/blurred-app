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

// Get the complete sidebar component (for initial load)
func (h *sidebarHandler) GetSidebar(c echo.Context) error {
	sportsList, err := h.sportService.GetAll()
	if err != nil {
		// Return sidebar with empty sports on error
		component := sidebar.SidebarWithSports([]sports.SportModel{})
		return component.Render(c.Request().Context(), c.Response().Writer)
	}

	component := sidebar.SidebarWithSports(sportsList)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

// Refresh just the sports list portion (for updates)
func (h *sidebarHandler) RefreshSports(c echo.Context) error {
	sports, err := h.sportService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get sports")
	}

	component := sidebar.SportsList(sports)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
