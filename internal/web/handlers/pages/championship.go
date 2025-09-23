package pages

import (
	"net/http"

	"github.com/JulOuellet/sportlight/internal/domains/championships"
	"github.com/JulOuellet/sportlight/templates/pages"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ChampionshipPageHandler struct {
	championService championships.ChampionshipService
}

func NewChampionshipPageHandler(
	championService championships.ChampionshipService,
) *ChampionshipPageHandler {
	return &ChampionshipPageHandler{
		championService: championService,
	}
}

func (h *ChampionshipPageHandler) GetChampionship(c echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid championship ID")
	}

	championship, err := h.championService.GetById(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Championship not found")
	}

	return pages.ChampionshipPage(championship).Render(c.Request().Context(), c.Response().Writer)
}
