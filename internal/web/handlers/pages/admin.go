package pages

import (
	"net/http"
	"os"

	"github.com/JulOuellet/blurred-app/internal/domains/championships"
	"github.com/JulOuellet/blurred-app/internal/domains/integrations"
	"github.com/JulOuellet/blurred-app/templates/admin"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AdminPageHandler struct {
	integrationService  integrations.IntegrationService
	championshipService championships.ChampionshipService
}

func NewAdminPageHandler(
	integrationService integrations.IntegrationService,
	championshipService championships.ChampionshipService,
) AdminPageHandler {
	return AdminPageHandler{
		integrationService:  integrationService,
		championshipService: championshipService,
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
	items, err := h.integrationService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load integrations")
	}
	if items == nil {
		items = []integrations.IntegrationModel{}
	}

	return admin.IntegrationsListPage(items).Render(c.Request().Context(), c.Response().Writer)
}

func (h *AdminPageHandler) NewIntegrationForm(c echo.Context) error {
	champs, err := h.championshipService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load championships")
	}

	return admin.IntegrationFormPage(admin.IntegrationFormData{
		Championships: champs,
	}).Render(c.Request().Context(), c.Response().Writer)
}

func (h *AdminPageHandler) CreateIntegration(c echo.Context) error {
	championshipID, err := uuid.Parse(c.FormValue("championshipId"))
	if err != nil {
		return h.renderFormWithError(c, "Invalid championship ID")
	}

	req := integrations.IntegrationRequest{
		YoutubeChannelID:   c.FormValue("youtubeChannelId"),
		YoutubeChannelName: c.FormValue("youtubeChannelName"),
		ChampionshipID:     championshipID,
		Lang:               c.FormValue("lang"),
		RelevancePattern:   c.FormValue("relevancePattern"),
		EventPattern:       c.FormValue("eventPattern"),
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
	champs, err := h.championshipService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load championships")
	}

	return admin.IntegrationFormPage(admin.IntegrationFormData{
		Championships: champs,
		ErrorMsg:      errorMsg,
	}).Render(c.Request().Context(), c.Response().Writer)
}
