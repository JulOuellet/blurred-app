package inbox

import (
	"database/sql"
	"fmt"
	"log"
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

	log.Printf("[processor] processing video %s: %s", item.YoutubeVideoID, item.VideoTitle)

	integration, err := p.integrationRepo.GetById(item.IntegrationID)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("integration not found: %v", err))
	}

	relevanceRe, err := regexp.Compile(integration.RelevancePattern)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("invalid relevance pattern: %v", err))
	}

	if !youtube.MatchesRelevancePattern(item.VideoTitle, relevanceRe) {
		log.Printf("[processor] skipping video %s: not relevant", item.YoutubeVideoID)
		return true, p.inboxRepo.MarkSkipped(item.ID, "not relevant")
	}

	eventRe, err := regexp.Compile(integration.EventPattern)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("invalid event pattern: %v", err))
	}

	eventNumber, ok := youtube.ExtractEventNumber(item.VideoTitle, eventRe)
	if !ok {
		return true, p.fail(item, "no event number found in title")
	}

	championshipEvents, err := p.eventRepo.GetAllByChampionshipId(
		integration.ChampionshipID,
		events.SortByDate,
		events.SortDirectionAsc,
	)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("failed to get events: %v", err))
	}

	matchedEvent := matchEventByNumber(championshipEvents, eventNumber)
	if matchedEvent == nil {
		return true, p.fail(item, fmt.Sprintf("no event matching number %d", eventNumber))
	}

	exists, err := p.highlightRepo.ExistsByYoutubeID(item.YoutubeVideoID)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("failed to check for duplicate: %v", err))
	}
	if exists {
		log.Printf("[processor] skipping video %s: duplicate", item.YoutubeVideoID)
		return true, p.inboxRepo.MarkSkipped(item.ID, "duplicate")
	}

	videoDetails, err := p.youtubeClient.GetVideoDetails(item.YoutubeVideoID)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("YouTube API error: %v", err))
	}

	videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.YoutubeVideoID)
	genericName := fmt.Sprintf("Stage %d Highlights", eventNumber)
	source := ""
	if integration.Source != nil {
		source = *integration.Source
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

	log.Printf("[processor] created highlight for video %s → event %s", item.YoutubeVideoID, matchedEvent.Name)
	return true, p.inboxRepo.MarkCompleted(item.ID)
}

func (p *Processor) fail(item *InboxItem, reason string) error {
	log.Printf("[processor] failed video %s: %s", item.YoutubeVideoID, reason)
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
