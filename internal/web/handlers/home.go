package handlers

import (
	"github.com/JulOuellet/sportlight/templates/pages"
	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context) error {
	return pages.Home().Render(c.Request().Context(), c.Response().Writer)
}
