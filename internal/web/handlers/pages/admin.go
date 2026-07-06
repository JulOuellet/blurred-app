package pages

import (
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/JulOuellet/blurred-app/internal/domains/highlights"
	"github.com/JulOuellet/blurred-app/internal/domains/integrations"
	"github.com/JulOuellet/blurred-app/internal/domains/sports"
	"github.com/JulOuellet/blurred-app/internal/inbox"
	"github.com/JulOuellet/blurred-app/templates/admin"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AdminPageHandler struct {
	integrationService integrations.IntegrationService
	sportService       sports.SportService
	inboxRepo          inbox.InboxRepository
	highlightRepo      highlights.HighlightRepository
}

func NewAdminPageHandler(
	integrationService integrations.IntegrationService,
	sportService sports.SportService,
	inboxRepo inbox.InboxRepository,
	highlightRepo highlights.HighlightRepository,
) AdminPageHandler {
	return AdminPageHandler{
		integrationService: integrationService,
		sportService:       sportService,
		inboxRepo:          inboxRepo,
		highlightRepo:      highlightRepo,
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
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int((7 * 24 * time.Hour).Seconds()),
	})

	return c.Redirect(http.StatusFound, "/admin/integrations")
}

func (h *AdminPageHandler) PostLogout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "admin_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
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
	req, err := integrationRequestFromForm(c)
	if err != nil {
		return h.renderFormWithError(c, "Invalid sport ID", nil)
	}

	if _, err := h.integrationService.Create(req); err != nil {
		return h.renderFormWithError(c, err.Error(), nil)
	}

	return c.Redirect(http.StatusFound, "/admin/integrations")
}

func (h *AdminPageHandler) GetIntegration(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid integration ID")
	}

	integration, err := h.integrationService.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Integration not found")
	}

	allSports, err := h.sportService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load sports")
	}

	return admin.IntegrationFormPage(admin.IntegrationFormData{
		Sports:      allSports,
		Integration: integration,
	}).Render(c.Request().Context(), c.Response().Writer)
}

func (h *AdminPageHandler) UpdateIntegration(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid integration ID")
	}

	existing, err := h.integrationService.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Integration not found")
	}

	req, err := integrationRequestFromForm(c)
	if err != nil {
		return h.renderFormWithError(c, "Invalid sport ID", existing)
	}

	if _, err := h.integrationService.Update(id, req); err != nil {
		// Re-render with the submitted values so a regex typo doesn't
		// wipe the rest of the edits.
		return h.renderFormWithError(c, err.Error(), modelFromRequest(existing, req))
	}

	return c.Redirect(http.StatusFound, "/admin/integrations")
}

func integrationRequestFromForm(c echo.Context) (integrations.IntegrationRequest, error) {
	sportID, err := uuid.Parse(c.FormValue("sportId"))
	if err != nil {
		return integrations.IntegrationRequest{}, err
	}

	return integrations.IntegrationRequest{
		YoutubeChannelID:   c.FormValue("youtubeChannelId"),
		YoutubeChannelName: c.FormValue("youtubeChannelName"),
		SportID:            sportID,
		Lang:               c.FormValue("lang"),
		ContentFilter:      c.FormValue("contentFilter"),
		TitleExclude:       c.FormValue("titleExclude"),
		StagePattern:       c.FormValue("stagePattern"),
		Active:             c.FormValue("active") == "on",
	}, nil
}

func modelFromRequest(existing *integrations.IntegrationModel, req integrations.IntegrationRequest) *integrations.IntegrationModel {
	m := *existing
	m.YoutubeChannelID = req.YoutubeChannelID
	m.YoutubeChannelName = optional(req.YoutubeChannelName)
	m.SportID = req.SportID
	m.Lang = req.Lang
	m.ContentFilter = optional(req.ContentFilter)
	m.TitleExclude = optional(req.TitleExclude)
	m.StagePattern = optional(req.StagePattern)
	m.Active = req.Active
	return &m
}

func optional(s string) *string {
	if s == "" {
		return nil
	}
	return &s
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

func (h *AdminPageHandler) ListInbox(c echo.Context) error {
	status := c.QueryParam("status")
	if status != "" && !slices.Contains(inbox.AllStatuses, status) {
		status = ""
	}

	items, err := h.inboxRepo.List(status, 100)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load inbox")
	}

	counts, err := h.inboxRepo.CountsByStatus()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load inbox counts")
	}

	return admin.InboxListPage(admin.InboxListData{
		Items:        items,
		Counts:       counts,
		ActiveFilter: status,
	}).Render(c.Request().Context(), c.Response().Writer)
}

func (h *AdminPageHandler) RemoveHighlight(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid inbox item ID")
	}

	item, err := h.inboxRepo.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Inbox item not found")
	}

	if err := h.highlightRepo.DeleteByYoutubeID(item.YoutubeVideoID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete highlight")
	}

	if err := h.inboxRepo.MarkSkipped(item.ID, "highlight removed by admin"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update inbox item")
	}

	redirect := "/admin/inbox"
	if status := c.FormValue("status"); status != "" {
		redirect += "?status=" + status
	}
	return c.Redirect(http.StatusFound, redirect)
}

func (h *AdminPageHandler) RetryInboxItem(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid inbox item ID")
	}

	if err := h.inboxRepo.Retry(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retry inbox item")
	}

	redirect := "/admin/inbox"
	if status := c.FormValue("status"); status != "" {
		redirect += "?status=" + status
	}
	return c.Redirect(http.StatusFound, redirect)
}

func (h *AdminPageHandler) renderFormWithError(c echo.Context, errorMsg string, integration *integrations.IntegrationModel) error {
	allSports, err := h.sportService.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load sports")
	}

	return admin.IntegrationFormPage(admin.IntegrationFormData{
		Sports:      allSports,
		ErrorMsg:    errorMsg,
		Integration: integration,
	}).Render(c.Request().Context(), c.Response().Writer)
}
