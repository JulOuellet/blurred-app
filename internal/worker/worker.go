package worker

import (
	"context"
	"log"
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
	log.Println("[worker] starting poller and processor")
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
			log.Println("[poller] shutting down")
			return
		case <-ticker.C:
			w.poll()
		}
	}
}

func (w *Worker) poll() {
	activeIntegrations, err := w.integrationRepo.GetAllActive()
	if err != nil {
		log.Printf("[poller] failed to get active integrations: %v", err)
		return
	}

	if len(activeIntegrations) == 0 {
		return
	}

	log.Printf("[poller] polling %d active integration(s)", len(activeIntegrations))

	for _, integration := range activeIntegrations {
		entries, err := youtube.FetchFeed(integration.YoutubeChannelID)
		if err != nil {
			log.Printf("[poller] failed to fetch feed for channel %s: %v", integration.YoutubeChannelID, err)
			continue
		}

		for _, entry := range entries {
			if err := w.inboxRepo.Insert(integration.ID, entry.VideoID, entry.Title, entry.PublishedAt); err != nil {
				log.Printf("[poller] failed to insert inbox entry for video %s: %v", entry.VideoID, err)
			}
		}

		now := time.Now()
		if err := w.integrationRepo.UpdateLastPolledAt(integration.ID, now); err != nil {
			log.Printf("[poller] failed to update last_polled_at for integration %s: %v", integration.ID, err)
		}
	}
}

func (w *Worker) runProcessor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("[processor] shutting down")
			return
		default:
			processed, err := w.processor.ProcessNext()
			if err != nil {
				log.Printf("[processor] error: %v", err)
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
