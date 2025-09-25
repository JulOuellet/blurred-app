package web

import (
	"github.com/JulOuellet/sportlight/internal/domains/championships"
	"github.com/JulOuellet/sportlight/internal/domains/seasons"
	"github.com/JulOuellet/sportlight/internal/domains/sports"
	"github.com/JulOuellet/sportlight/internal/web/handlers/components/sidebar"
	"github.com/JulOuellet/sportlight/internal/web/handlers/pages"
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

	sportRepository := sports.NewSportRepository(db)
	seasonRepository := seasons.NewSeasonRepository(db)
	championshipRepository := championships.NewChampionshipRepository(db)

	sportService := sports.NewSportService(sportRepository, seasonRepository)
	seasonService := seasons.NewSeasonService(seasonRepository)
	championshipService := championships.NewChampionshipService(championshipRepository)

	api := e.Group("/api")
	{
		sportHandler := sports.NewSportHandler(sportService)
		sportsApi := api.Group("/sports")
		sportsApi.GET("", sportHandler.GetAll)
		sportsApi.GET("/:id", sportHandler.GetById)
		sportsApi.POST("", sportHandler.Create)
		sportsApi.PATCH("/:id", sportHandler.Update)

		seasonHandler := seasons.NewSeasonHandler(seasonService)
		seasonsApi := api.Group("/seasons")
		seasonsApi.GET("", seasonHandler.GetAll)
		seasonsApi.GET("/:id", seasonHandler.GetById)
		seasonsApi.POST("", seasonHandler.Create)

		championshipHandler := championships.NewChampionshipHandler(championshipService)
		championshipsApi := api.Group("/championships")
		championshipsApi.GET("", championshipHandler.GetAll)
		championshipsApi.GET("/:id", championshipHandler.GetById)
		championshipsApi.POST("", championshipHandler.Create)
	}

	components := e.Group("/components")
	{
		sidebarHandler := sidebar.NewSidebarHandler(sportService)

		sidebarRoutes := components.Group("/sidebar")
		sidebarRoutes.GET("", sidebarHandler.GetSidebar)
		sidebarRoutes.GET("/sports", sidebarHandler.RefreshSports)
	}

	seasonPageHandler := pages.NewSeasonPageHandler(
		seasonService,
		championshipService,
		sportService,
	)
	e.GET("/", pages.HomePage)
	e.GET("/seasons/:id", seasonPageHandler.GetSeason)

	championshipPageHandler := pages.NewChampionshipPageHandler(
		championshipService,
		seasonService,
		sportService,
	)
	e.GET("/championships/:id", championshipPageHandler.GetChampionship)

	return e
}
