package inbox

import (
	"database/sql"
	"fmt"
	"log/slog"
	"regexp"
	"strconv"

	"github.com/JulOuellet/blurred-app/internal/domains/championships"
	"github.com/JulOuellet/blurred-app/internal/domains/events"
	"github.com/JulOuellet/blurred-app/internal/domains/highlights"
	"github.com/JulOuellet/blurred-app/internal/domains/integrations"
	"github.com/JulOuellet/blurred-app/internal/youtube"
)

type Processor struct {
	inboxRepo        InboxRepository
	integrationRepo  integrations.IntegrationRepository
	championshipRepo championships.ChampionshipRepository
	eventRepo        events.EventRepository
	highlightRepo    highlights.HighlightRepository
	youtubeClient    *youtube.Client
}

func NewProcessor(
	inboxRepo InboxRepository,
	integrationRepo integrations.IntegrationRepository,
	championshipRepo championships.ChampionshipRepository,
	eventRepo events.EventRepository,
	highlightRepo highlights.HighlightRepository,
	youtubeClient *youtube.Client,
) *Processor {
	return &Processor{
		inboxRepo:        inboxRepo,
		integrationRepo:  integrationRepo,
		championshipRepo: championshipRepo,
		eventRepo:        eventRepo,
		highlightRepo:    highlightRepo,
		youtubeClient:    youtubeClient,
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

	// Step 1: Content filter — skip non-highlight videos
	if integration.ContentFilter != nil {
		contentRe, err := regexp.Compile(*integration.ContentFilter)
		if err != nil {
			return true, p.fail(item, fmt.Sprintf("invalid content filter: %v", err))
		}
		if !contentRe.MatchString(item.VideoTitle) {
			slog.Debug("skipping video: not relevant", "component", "processor", "video_id", item.YoutubeVideoID)
			return true, p.inboxRepo.MarkSkipped(item.ID, "not relevant")
		}
	}

	// Step 1b: Title exclude — skip videos matching the exclude pattern
	if integration.TitleExclude != nil {
		excludeRe, err := regexp.Compile(*integration.TitleExclude)
		if err != nil {
			return true, p.fail(item, fmt.Sprintf("invalid title exclude pattern: %v", err))
		}
		if excludeRe.MatchString(item.VideoTitle) {
			slog.Debug("skipping video: matched title exclude", "component", "processor", "video_id", item.YoutubeVideoID)
			return true, p.inboxRepo.MarkSkipped(item.ID, "excluded by title exclude pattern")
		}
	}

	// Step 2: Championship matching — find which championship this video belongs to
	if item.PublishedAt == nil {
		return true, p.fail(item, "video has no published_at date")
	}

	candidates, err := p.championshipRepo.GetCandidatesForMatching(integration.SportID, *item.PublishedAt)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("failed to get candidate championships: %v", err))
	}

	var matchedChampionship *championships.ChampionshipModel
	for i := range candidates {
		if candidates[i].TitlePattern == nil {
			continue
		}
		titleRe, err := regexp.Compile(*candidates[i].TitlePattern)
		if err != nil {
			slog.Warn("invalid title pattern", "component", "processor", "championship", candidates[i].Name, "error", err)
			continue
		}
		if titleRe.MatchString(item.VideoTitle) {
			matchedChampionship = &candidates[i]
			break
		}
	}

	if matchedChampionship == nil {
		slog.Debug("skipping video: no matching championship", "component", "processor", "video_id", item.YoutubeVideoID)
		return true, p.inboxRepo.MarkSkipped(item.ID, "no matching championship")
	}

	// Step 3: Event matching — find the right event within the championship
	championshipEvents, err := p.eventRepo.GetAllByChampionshipId(
		matchedChampionship.ID,
		events.SortByDate,
		events.SortDirectionAsc,
	)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("failed to get events: %v", err))
	}

	if len(championshipEvents) == 0 {
		return true, p.fail(item, fmt.Sprintf("no events for championship %q", matchedChampionship.Name))
	}

	var matchedEvent *events.EventModel
	if len(championshipEvents) == 1 {
		matchedEvent = &championshipEvents[0]
	} else if integration.StagePattern == nil {
		return true, p.fail(item, "stage pattern required for championships with multiple events")
	} else {
		stageRe, err := regexp.Compile(*integration.StagePattern)
		if err != nil {
			return true, p.fail(item, fmt.Sprintf("invalid stage pattern: %v", err))
		}

		eventNumber, ok := youtube.ExtractEventNumber(item.VideoTitle, stageRe)
		if !ok {
			return true, p.fail(item, "no stage number found in title")
		}

		matchedEvent = matchEventByNumber(championshipEvents, eventNumber)
		if matchedEvent == nil {
			return true, p.fail(item, fmt.Sprintf("no event matching stage %d", eventNumber))
		}
	}

	// Step 4: Duplicate check
	exists, err := p.highlightRepo.ExistsByYoutubeID(item.YoutubeVideoID)
	if err != nil {
		return true, p.fail(item, fmt.Sprintf("failed to check for duplicate: %v", err))
	}
	if exists {
		slog.Debug("skipping video: duplicate", "component", "processor", "video_id", item.YoutubeVideoID)
		return true, p.inboxRepo.MarkSkipped(item.ID, "duplicate")
	}

	// Step 5: Fetch video details and create highlight
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

	slog.Info("created highlight",
		"component", "processor",
		"video_id", item.YoutubeVideoID,
		"championship", matchedChampionship.Name,
		"event", matchedEvent.Name,
	)
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
