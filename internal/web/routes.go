package web

import (
	"net/http"

	"github.com/JulOuellet/sportlight/internal/domains/sports"
	"github.com/JulOuellet/sportlight/internal/pages"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(db *sqlx.DB) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	api := e.Group("/api")

	sportRepository := sports.NewSportRepository(db)
	sportService := sports.NewSportService(sportRepository)
	sportHandler := sports.NewSportHandler(sportService)

	sports := api.Group("/sports")
	sports.GET("", sportHandler.GetAll)
	sports.GET("/:id", sportHandler.GetById)
	sports.POST("", sportHandler.Create)
	sports.PATCH("/:id", sportHandler.Update)

	// seasonsRepository := seasons.NewSeasonRepository(db)
	// seasonsService := seasons.NewSeasonService(seasonsRepository)
	// seasonsHandler := seasons.NewSeasonHandler(seasonsService)

	// e.GET("/seasons", seasonsHandler.GetAll)

	e.GET("/", echo.WrapHandler(http.HandlerFunc(pages.IndexHandler)))

	return e
}
