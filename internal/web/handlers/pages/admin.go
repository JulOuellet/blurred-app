package pages

import (
	"net/http"
	"os"

	"github.com/JulOuellet/blurred-app/internal/domains/integrations"
	"github.com/JulOuellet/blurred-app/internal/domains/sports"
	"github.com/JulOuellet/blurred-app/templates/admin"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AdminPageHandler struct {
	integrationService integrations.IntegrationService
	sportService       sports.SportService
}

func NewAdminPageHandler(
	integrationService integrations.IntegrationService,
	sportService sports.SportService,
) AdminPageHandler {
	return AdminPageHandler{
		integrationService: integrationService,
		sportService:       sportService,
	}
}

func (h *AdminPageHandler) GetLogin(c echo.Context) error {
	return admin.LoginPage("").Render(c.Request().Context(), c.Response().Writer)
}

func (h *AdminPageHandler) PostLogin(c echo.Context) error {
	password := c.FormValue("password")
	hash := os.Getenv("ADMIN_PASSWORD_HASH")
	if hash == "" || bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return admin.LoginPage("Invalid password").Render(c.Request().Context(), c.Response().Writer)
	}

	secret := os.Getenv("ADMIN_SESSION_SECRET")
	c.SetCookie(&http.Cookie{
		Name:     "admin_token",
		Value:    secret,
		Path:     "/admin",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	return c.Redirect(http.StatusFound, "/admin/integrations")
}

func (h *AdminPageHandler) PostLogout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "admin_token",
		Value:    "",
		Path:     "/admin",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	})

	return c.Redirect(http.StatusFound, "/admin/login")
}

func (h *AdminPageHandler) ListIntegrations(c echo.Context) error {
	items, err := h.integrationService.GetAllWithSport()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load integrations")
	}
	if items == nil {
		items = []integrations.IntegrationWithSport{}
	}

	allSports, err := h.sportService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load sports")
	}

	filter := c.QueryParam("sport")
	if filter != "" {
		filterID, err := uuid.Parse(filter)
		if err == nil {
			filtered := []integrations.IntegrationWithSport{}
			for _, item := range items {
				if item.SportID == filterID {
					filtered = append(filtered, item)
				}
			}
			items = filtered
		}
	}

	return admin.IntegrationsListPage(admin.IntegrationsListData{
		Items:        items,
		Sports:       allSports,
		ActiveFilter: filter,
	}).Render(c.Request().Context(), c.Response().Writer)
}

func (h *AdminPageHandler) NewIntegrationForm(c echo.Context) error {
	allSports, err := h.sportService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load sports")
	}

	return admin.IntegrationFormPage(admin.IntegrationFormData{
		Sports: allSports,
	}).Render(c.Request().Context(), c.Response().Writer)
}

func (h *AdminPageHandler) CreateIntegration(c echo.Context) error {
	sportID, err := uuid.Parse(c.FormValue("sportId"))
	if err != nil {
		return h.renderFormWithError(c, "Invalid sport ID")
	}

	req := integrations.IntegrationRequest{
		YoutubeChannelID:   c.FormValue("youtubeChannelId"),
		YoutubeChannelName: c.FormValue("youtubeChannelName"),
		SportID:            sportID,
		Lang:               c.FormValue("lang"),
		ContentFilter:      c.FormValue("contentFilter"),
		TitleExclude:       c.FormValue("titleExclude"),
		StagePattern:       c.FormValue("stagePattern"),
	}

	_, err = h.integrationService.Create(req)
	if err != nil {
		return h.renderFormWithError(c, err.Error())
	}

	return c.Redirect(http.StatusFound, "/admin/integrations")
}

func (h *AdminPageHandler) DeleteIntegration(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid integration ID")
	}

	if err := h.integrationService.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete integration")
	}

	return c.Redirect(http.StatusFound, "/admin/integrations")
}

func (h *AdminPageHandler) renderFormWithError(c echo.Context, errorMsg string) error {
	allSports, err := h.sportService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load sports")
	}

	return admin.IntegrationFormPage(admin.IntegrationFormData{
		Sports:   allSports,
		ErrorMsg: errorMsg,
	}).Render(c.Request().Context(), c.Response().Writer)
}
