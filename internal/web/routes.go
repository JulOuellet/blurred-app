package web

import (
	"github.com/JulOuellet/sportlight/internal/domains/sports"
	"github.com/JulOuellet/sportlight/internal/web/handlers"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(db *sqlx.DB) *echo.Echo {
	e := echo.New()

	e.Static("/assets", "assets")

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

	e.GET("/", handlers.HomePage)
	return e
}
