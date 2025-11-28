package web

import (
	"github.com/JulOuellet/blurred-app/internal/domains/championships"
	"github.com/JulOuellet/blurred-app/internal/domains/events"
	"github.com/JulOuellet/blurred-app/internal/domains/highlights"
	"github.com/JulOuellet/blurred-app/internal/domains/seasons"
	"github.com/JulOuellet/blurred-app/internal/domains/sports"
	"github.com/JulOuellet/blurred-app/internal/web/handlers/pages"
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
	eventRepository := events.NewEventRepository(db)
	highlightRepository := highlights.NewHighlightRepository(db)

	sportService := sports.NewSportService(sportRepository, seasonRepository)
	seasonService := seasons.NewSeasonService(seasonRepository)
	championshipService := championships.NewChampionshipService(championshipRepository)
	eventService := events.NewEventService(eventRepository)
	highlightService := highlights.NewHighlightService(highlightRepository)

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

		eventHandler := events.NewEventHandler(eventService)
		eventsApi := api.Group("/events")
		eventsApi.GET("", eventHandler.GetAll)
		eventsApi.GET("/:id", eventHandler.GetById)
		eventsApi.POST("", eventHandler.Create)
		eventsApi.GET("/championship/:id", eventHandler.GetAllByChampionshipId)

		highlightHandler := highlights.NewHighlightHandler(highlightService)
		highlightsApi := api.Group("/highlights")
		highlightsApi.GET("", highlightHandler.GetAll)
		highlightsApi.GET("/:id", highlightHandler.GetById)
		highlightsApi.POST("", highlightHandler.Create)
		highlightsApi.GET("/event/:id", highlightHandler.GetAllByEventId)
	}

	seasonPageHandler := pages.NewSeasonPageHandler(
		seasonService,
		championshipService,
		sportService,
	)
	homePageHandler := pages.NewHomePageHandler(sportService)
	e.GET("/", homePageHandler.GetHome)
	e.GET("/seasons/:id", seasonPageHandler.GetSeason)

	championshipPageHandler := pages.NewChampionshipPageHandler(
		championshipService,
		seasonService,
		sportService,
		eventService,
	)
	e.GET("/championships/:id", championshipPageHandler.GetChampionship)

	eventPageHandler := pages.NewEventPageHandler(
		championshipService,
		seasonService,
		sportService,
		eventService,
		highlightService,
	)
	e.GET("/events/:id", eventPageHandler.GetEvent)

	return e
}
