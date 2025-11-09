package highlights

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type HighlightService interface {
	GetAll() ([]HighlightModel, error)
	GetById(id uuid.UUID) (*HighlightModel, error)
	Create(req HighlightRequest) (*HighlightModel, error)
	GetAllByEventId(eventId uuid.UUID) ([]HighlightModel, error)
}

type highlightService struct {
	highlightRepo HighlightRepository
}

func NewHighlightService(highlightRepo HighlightRepository) HighlightService {
	return &highlightService{highlightRepo: highlightRepo}
}

func (s *highlightService) GetAll() ([]HighlightModel, error) {
	return s.highlightRepo.GetAll()
}

func (s *highlightService) GetById(id uuid.UUID) (*HighlightModel, error) {
	return s.highlightRepo.GetById(id)
}

func (s *highlightService) GetAllByEventId(eventId uuid.UUID) ([]HighlightModel, error) {
	return s.highlightRepo.GetAllByEventId(eventId)
}

func (s *highlightService) Create(req HighlightRequest) (*HighlightModel, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, fmt.Errorf("highlight name cannot be empty")
	}

	genericName := strings.TrimSpace(req.GenericName)
	if genericName == "" {
		return nil, fmt.Errorf("highlight generic name cannot be empty")
	}

	url := strings.TrimSpace(req.URL)
	if url == "" {
		return nil, fmt.Errorf("highlight url cannot be empty")
	}

	youtubeID := strings.TrimSpace(req.YoutubeID)

	language := strings.TrimSpace(req.Language)
	if language == "" {
		return nil, fmt.Errorf("highlight language cannot be empty")
	}

	mediaType := strings.TrimSpace(req.MediaType)
	if mediaType == "" {
		return nil, fmt.Errorf("highlight media type cannot be empty")
	}

	source := strings.TrimSpace(req.Source)

	return s.highlightRepo.Create(
		name,
		genericName,
		url,
		youtubeID,
		req.DurationSeconds,
		language,
		mediaType,
		source,
		req.EventID,
	)
}
