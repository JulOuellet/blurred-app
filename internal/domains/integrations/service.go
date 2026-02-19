package integrations

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/JulOuellet/blurred-app/internal/domains/highlights"
	"github.com/google/uuid"
)

type IntegrationService interface {
	GetAll() ([]IntegrationModel, error)
	GetById(id uuid.UUID) (*IntegrationModel, error)
	GetAllActive() ([]IntegrationModel, error)
	Create(req IntegrationRequest) (*IntegrationModel, error)
}

type integrationService struct {
	integrationRepo IntegrationRepository
}

func NewIntegrationService(integrationRepo IntegrationRepository) IntegrationService {
	return &integrationService{integrationRepo: integrationRepo}
}

func (s *integrationService) GetAll() ([]IntegrationModel, error) {
	return s.integrationRepo.GetAll()
}

func (s *integrationService) GetById(id uuid.UUID) (*IntegrationModel, error) {
	return s.integrationRepo.GetById(id)
}

func (s *integrationService) GetAllActive() ([]IntegrationModel, error) {
	return s.integrationRepo.GetAllActive()
}

func (s *integrationService) Create(req IntegrationRequest) (*IntegrationModel, error) {
	youtubeChannelID := strings.TrimSpace(req.YoutubeChannelID)
	if youtubeChannelID == "" {
		return nil, fmt.Errorf("youtube channel ID cannot be empty")
	}

	lang := strings.TrimSpace(req.Lang)
	if lang == "" {
		return nil, fmt.Errorf("language cannot be empty")
	}
	if !highlights.IsValidLanguage(lang) {
		return nil, fmt.Errorf("invalid language %q: must be a valid language code (e.g. en-GB, fr-FR)", lang)
	}

	relevancePattern := strings.TrimSpace(req.RelevancePattern)
	if relevancePattern == "" {
		return nil, fmt.Errorf("relevance pattern cannot be empty")
	}
	if _, err := regexp.Compile(relevancePattern); err != nil {
		return nil, fmt.Errorf("invalid relevance pattern: %w", err)
	}

	eventPattern := strings.TrimSpace(req.EventPattern)
	if eventPattern == "" {
		return nil, fmt.Errorf("event pattern cannot be empty")
	}
	compiledEvent, err := regexp.Compile(eventPattern)
	if err != nil {
		return nil, fmt.Errorf("invalid event pattern: %w", err)
	}
	if compiledEvent.NumSubexp() < 1 {
		return nil, fmt.Errorf("event pattern must contain at least one capture group for the event number")
	}

	youtubeChannelName := strings.TrimSpace(req.YoutubeChannelName)
	source := strings.TrimSpace(req.Source)

	return s.integrationRepo.Create(
		youtubeChannelID,
		youtubeChannelName,
		req.ChampionshipID,
		lang,
		source,
		relevancePattern,
		eventPattern,
	)
}
