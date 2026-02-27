package web

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/JulOuellet/blurred-app/internal/domains/championships"
	"github.com/JulOuellet/blurred-app/internal/domains/events"
	"github.com/JulOuellet/blurred-app/internal/domains/highlights"
	"github.com/JulOuellet/blurred-app/internal/domains/integrations"
	"github.com/JulOuellet/blurred-app/internal/domains/search"
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
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogMethod:  true,
		LogLatency: true,
		LogError:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			attrs := []slog.Attr{
				slog.String("method", v.Method),
				slog.String("uri", v.URI),
				slog.Int("status", v.Status),
				slog.String("latency", v.Latency.String()),
			}
			if v.Error != nil {
				attrs = append(attrs, slog.String("error", v.Error.Error()))
				slog.LogAttrs(context.Background(), slog.LevelError, "request", attrs...)
			} else {
				slog.LogAttrs(context.Background(), slog.LevelInfo, "request", attrs...)
			}
			return nil
		},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	sportRepository := sports.NewSportRepository(db)
	seasonRepository := seasons.NewSeasonRepository(db)
	championshipRepository := championships.NewChampionshipRepository(db)
	eventRepository := events.NewEventRepository(db)
	highlightRepository := highlights.NewHighlightRepository(db)
	integrationRepository := integrations.NewIntegrationRepository(db)

	sportService := sports.NewSportService(sportRepository, seasonRepository)
	seasonService := seasons.NewSeasonService(seasonRepository)
	championshipService := championships.NewChampionshipService(championshipRepository)
	eventService := events.NewEventService(eventRepository)
	highlightService := highlights.NewHighlightService(highlightRepository)
	integrationService := integrations.NewIntegrationService(integrationRepository)

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

		integrationHandler := integrations.NewIntegrationHandler(integrationService)
		integrationsApi := api.Group("/integrations")
		integrationsApi.GET("", integrationHandler.GetAll)
		integrationsApi.GET("/:id", integrationHandler.GetById)
		integrationsApi.POST("", integrationHandler.Create)
	}

	seasonPageHandler := pages.NewSeasonPageHandler(
		seasonService,
		championshipService,
		sportService,
	)
	homePageHandler := pages.NewHomePageHandler(sportService, championshipService)
	e.GET("/", homePageHandler.GetHome)
	e.GET("/seasons/:id", seasonPageHandler.GetSeason)

	championshipPageHandler := pages.NewChampionshipPageHandler(
		championshipService,
		seasonService,
		sportService,
		eventService,
	)
	e.GET("/championships/:id", championshipPageHandler.GetChampionship)
	e.GET("/championships/:id/events", championshipPageHandler.GetChampionshipEvents)

	eventPageHandler := pages.NewEventPageHandler(
		championshipService,
		seasonService,
		sportService,
		eventService,
		highlightService,
	)
	e.GET("/events/:id", eventPageHandler.GetEvent)

	aboutPageHandler := pages.NewAboutPageHandler(sportService)
	e.GET("/about", aboutPageHandler.GetAbout)

	searchRepository := search.NewSearchRepository(db)
	searchService := search.NewSearchService(searchRepository)
	searchHandler := search.NewSearchHandler(searchService)
	e.GET("/search", searchHandler.Search)

	// Admin routes
	adminHandler := pages.NewAdminPageHandler(integrationService, championshipService)
	e.GET("/admin/login", adminHandler.GetLogin)
	e.POST("/admin/login", adminHandler.PostLogin)
	e.POST("/admin/logout", adminHandler.PostLogout)

	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "cookie:admin_token",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == os.Getenv("ADMIN_SESSION_SECRET"), nil
		},
		ErrorHandler: func(err error, c echo.Context) error {
			return c.Redirect(http.StatusFound, "/admin/login")
		},
	}))
	adminGroup.GET("/integrations", adminHandler.ListIntegrations)
	adminGroup.GET("/integrations/new", adminHandler.NewIntegrationForm)
	adminGroup.POST("/integrations", adminHandler.CreateIntegration)
	adminGroup.POST("/integrations/:id/delete", adminHandler.DeleteIntegration)

	return e
}
