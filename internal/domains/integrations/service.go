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
	GetAllWithSport() ([]IntegrationWithSport, error)
	GetById(id uuid.UUID) (*IntegrationModel, error)
	GetAllActive() ([]IntegrationModel, error)
	Create(req IntegrationRequest) (*IntegrationModel, error)
	Delete(id uuid.UUID) error
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

func (s *integrationService) GetAllWithSport() ([]IntegrationWithSport, error) {
	return s.integrationRepo.GetAllWithSport()
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

	if req.SportID == uuid.Nil {
		return nil, fmt.Errorf("sport ID cannot be empty")
	}

	lang := strings.TrimSpace(req.Lang)
	if lang == "" {
		return nil, fmt.Errorf("language cannot be empty")
	}
	if !highlights.IsValidLanguage(lang) {
		return nil, fmt.Errorf("invalid language %q: must be a valid language code (e.g. en-GB, fr-FR)", lang)
	}

	var contentFilter *string
	trimmedContentFilter := strings.TrimSpace(req.ContentFilter)
	if trimmedContentFilter != "" {
		if _, err := regexp.Compile(trimmedContentFilter); err != nil {
			return nil, fmt.Errorf("invalid content filter: %w", err)
		}
		contentFilter = &trimmedContentFilter
	}

	var titleExclude *string
	trimmedTitleExclude := strings.TrimSpace(req.TitleExclude)
	if trimmedTitleExclude != "" {
		if _, err := regexp.Compile(trimmedTitleExclude); err != nil {
			return nil, fmt.Errorf("invalid title exclude pattern: %w", err)
		}
		titleExclude = &trimmedTitleExclude
	}

	var stagePattern *string
	trimmedStagePattern := strings.TrimSpace(req.StagePattern)
	if trimmedStagePattern != "" {
		compiled, err := regexp.Compile(trimmedStagePattern)
		if err != nil {
			return nil, fmt.Errorf("invalid stage pattern: %w", err)
		}
		if compiled.NumSubexp() < 1 {
			return nil, fmt.Errorf("stage pattern must contain at least one capture group for the stage number")
		}
		stagePattern = &trimmedStagePattern
	}

	youtubeChannelName := strings.TrimSpace(req.YoutubeChannelName)

	return s.integrationRepo.Create(
		youtubeChannelID,
		youtubeChannelName,
		req.SportID,
		lang,
		contentFilter,
		titleExclude,
		stagePattern,
	)
}

func (s *integrationService) Delete(id uuid.UUID) error {
	return s.integrationRepo.Delete(id)
}
