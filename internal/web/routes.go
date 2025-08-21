package web

import (
	"database/sql"
	"net/http"

	"github.com/JulOuellet/sportlight/internal/domains/seasons"
	"github.com/JulOuellet/sportlight/internal/domains/sports"
	"github.com/JulOuellet/sportlight/internal/pages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(db *sql.DB) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	sportRepository := sports.NewSportRepository(db)
	sportService := sports.NewSportService(sportRepository)
	sportHandler := sports.NewSportHandler(sportService)

	e.GET("/sports", sportHandler.GetAll)

	seasonsRepository := seasons.NewSeasonRepository(db)
	seasonsService := seasons.NewSeasonService(seasonsRepository)
	seasonsHandler := seasons.NewSeasonHandler(seasonsService)

	e.GET("/seasons", seasonsHandler.GetAll)

	e.GET("/", echo.WrapHandler(http.HandlerFunc(pages.IndexHandler)))

	return e
}
