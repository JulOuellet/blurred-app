package inbox

import (
	"database/sql"
	"fmt"
	"log/slog"
	"regexp"
	"strconv"

	"github.com/JulOuellet/blurred-app/internal/domains/events"
	"github.com/JulOuellet/blurred-app/internal/domains/highlights"
	"github.com/JulOuellet/blurred-app/internal/domains/integrations"
	"github.com/JulOuellet/blurred-app/internal/youtube"
)

type Processor struct {
	inboxRepo       InboxRepository
	integrationRepo integrations.IntegrationRepository
	eventRepo       events.EventRepository
	highlightRepo   highlights.HighlightRepository
	youtubeClient   *youtube.Client
}

func NewProcessor(
	inboxRepo InboxRepository,
	integrationRepo integrations.IntegrationRepository,
	eventRepo events.EventRepository,
	highlightRepo highlights.HighlightRepository,
	youtubeClient *youtube.Client,
) *Processor {
	return &Processor{
		inboxRepo:       inboxRepo,
		integrationRepo: integrationRepo,
		eventRepo:       eventRepo,
		highlightRepo:   highlightRepo,
		youtubeClient:   youtubeClient,
	}
}

func (p *Processor) ProcessNext() (bool, error) {
	item, err := p.inboxRepo.ClaimNext()
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to claim inbox item: %w", err)
	}

	slog.Info("processing video", "component", "processor", "video_id", item.YoutubeVideoID, "title", item.VideoTitle)

	integration, err := p.integrationRepo.GetById(item.IntegrationID)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("integration not found: %v", err))
	}

	relevanceRe, err := regexp.Compile(integration.RelevancePattern)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("invalid relevance pattern: %v", err))
	}

	if !youtube.MatchesRelevancePattern(item.VideoTitle, relevanceRe) {
		slog.Debug("skipping video: not relevant", "component", "processor", "video_id", item.YoutubeVideoID)
		return true, p.inboxRepo.MarkSkipped(item.ID, "not relevant")
	}

	championshipEvents, err := p.eventRepo.GetAllByChampionshipId(
		integration.ChampionshipID,
		events.SortByDate,
		events.SortDirectionAsc,
	)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("failed to get events: %v", err))
	}

	var matchedEvent *events.EventModel
	if len(championshipEvents) == 1 {
		matchedEvent = &championshipEvents[0]
	} else if integration.EventPattern == nil {
		return true, p.fail(item, "event pattern is required for championships with multiple events")
	} else {
		eventRe, err := regexp.Compile(*integration.EventPattern)
		if err != nil {
			return true, p.fail(item, fmt.Sprintf("invalid event pattern: %v", err))
		}

		eventNumber, ok := youtube.ExtractEventNumber(item.VideoTitle, eventRe)
		if !ok {
			return true, p.fail(item, "no event number found in title")
		}

		matchedEvent = matchEventByNumber(championshipEvents, eventNumber)
		if matchedEvent == nil {
			return true, p.fail(item, fmt.Sprintf("no event matching number %d", eventNumber))
		}
	}

	exists, err := p.highlightRepo.ExistsByYoutubeID(item.YoutubeVideoID)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("failed to check for duplicate: %v", err))
	}
	if exists {
		slog.Debug("skipping video: duplicate", "component", "processor", "video_id", item.YoutubeVideoID)
		return true, p.inboxRepo.MarkSkipped(item.ID, "duplicate")
	}

	videoDetails, err := p.youtubeClient.GetVideoDetails(item.YoutubeVideoID)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("YouTube API error: %v", err))
	}

	videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.YoutubeVideoID)
	genericName := fmt.Sprintf("%s Highlights", matchedEvent.Name)
	source := ""
	if integration.YoutubeChannelName != nil {
		source = *integration.YoutubeChannelName
	}

	_, err = p.highlightRepo.Create(
		videoDetails.Title,
		genericName,
		videoURL,
		item.YoutubeVideoID,
		videoDetails.DurationSeconds,
		integration.Lang,
		"VIDEO",
		source,
		matchedEvent.ID,
	)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("failed to create highlight: %v", err))
	}

	slog.Info("created highlight", "component", "processor", "video_id", item.YoutubeVideoID, "event", matchedEvent.Name)
	return true, p.inboxRepo.MarkCompleted(item.ID)
}

func (p *Processor) fail(item *InboxItem, reason string) error {
	slog.Warn("failed video", "component", "processor", "video_id", item.YoutubeVideoID, "reason", reason)
	return p.inboxRepo.MarkFailed(item.ID, reason)
}

var stageNumberRe = regexp.MustCompile(`(?i)(?:stage|[éeè]tape)\s+(\d+)`)

func matchEventByNumber(evts []events.EventModel, number int) *events.EventModel {
	for i := range evts {
		matches := stageNumberRe.FindStringSubmatch(evts[i].Name)
		if len(matches) >= 2 {
			n, err := strconv.Atoi(matches[1])
			if err == nil && n == number {
				return &evts[i]
			}
		}
	}

	// Fallback: match by position (event at index N-1 = Stage N)
	if number >= 1 && number <= len(evts) {
		return &evts[number-1]
	}

	return nil
}
