package worker

import (
	"context"
	"log/slog"
	"time"

	"github.com/JulOuellet/blurred-app/internal/domains/events"
	"github.com/JulOuellet/blurred-app/internal/domains/highlights"
	"github.com/JulOuellet/blurred-app/internal/domains/integrations"
	"github.com/JulOuellet/blurred-app/internal/inbox"
	"github.com/JulOuellet/blurred-app/internal/youtube"
	"github.com/jmoiron/sqlx"
)

const (
	pollInterval     = 5 * time.Minute
	processIdleWait  = 10 * time.Second
	processErrorWait = 30 * time.Second
)

type Worker struct {
	integrationRepo integrations.IntegrationRepository
	inboxRepo       inbox.InboxRepository
	processor       *inbox.Processor
}

func New(db *sqlx.DB, youtubeAPIKey string) *Worker {
	integrationRepo := integrations.NewIntegrationRepository(db)
	inboxRepo := inbox.NewInboxRepository(db)
	eventRepo := events.NewEventRepository(db)
	highlightRepo := highlights.NewHighlightRepository(db)
	youtubeClient := youtube.NewClient(youtubeAPIKey)

	processor := inbox.NewProcessor(
		inboxRepo,
		integrationRepo,
		eventRepo,
		highlightRepo,
		youtubeClient,
	)

	return &Worker{
		integrationRepo: integrationRepo,
		inboxRepo:       inboxRepo,
		processor:       processor,
	}
}

func (w *Worker) Run(ctx context.Context) {
	slog.Info("starting poller and processor", "component", "worker")
	go w.runPoller(ctx)
	go w.runProcessor(ctx)
}

func (w *Worker) runPoller(ctx context.Context) {
	// Run immediately on startup, then every pollInterval
	w.poll()

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			slog.Info("shutting down", "component", "poller")
			return
		case <-ticker.C:
			w.poll()
		}
	}
}

func (w *Worker) poll() {
	activeIntegrations, err := w.integrationRepo.GetAllActive()
	if err != nil {
		slog.Error("failed to get active integrations", "component", "poller", "error", err)
		return
	}

	if len(activeIntegrations) == 0 {
		return
	}

	slog.Info("polling active integrations", "component", "poller", "count", len(activeIntegrations))

	for _, integration := range activeIntegrations {
		entries, err := youtube.FetchFeed(integration.YoutubeChannelID)
		if err != nil {
			slog.Error("failed to fetch feed", "component", "poller", "channel", integration.YoutubeChannelID, "error", err)
			continue
		}

		for _, entry := range entries {
			if err := w.inboxRepo.Insert(integration.ID, entry.VideoID, entry.Title, entry.PublishedAt); err != nil {
				slog.Error("failed to insert inbox entry", "component", "poller", "video_id", entry.VideoID, "error", err)
			}
		}

		now := time.Now()
		if err := w.integrationRepo.UpdateLastPolledAt(integration.ID, now); err != nil {
			slog.Error("failed to update last_polled_at", "component", "poller", "integration_id", integration.ID, "error", err)
		}
	}
}

func (w *Worker) runProcessor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			slog.Info("shutting down", "component", "processor")
			return
		default:
			processed, err := w.processor.ProcessNext()
			if err != nil {
				slog.Error("processing error", "component", "processor", "error", err)
				sleep(ctx, processErrorWait)
				continue
			}
			if !processed {
				sleep(ctx, processIdleWait)
			}
		}
	}
}

func sleep(ctx context.Context, d time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(d):
	}
}
