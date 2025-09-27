package events

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type EventService interface {
	GetAll() ([]EventModel, error)
	GetById(id uuid.UUID) (*EventModel, error)
	Create(req EventRequest) (*EventModel, error)
	GetAllByChampionshipId(championshipId uuid.UUID) ([]EventModel, error)
}

type eventService struct {
	eventRepo EventRepository
}

func NewEventService(eventRepo EventRepository) EventService {
	return &eventService{eventRepo: eventRepo}
}

func (s *eventService) GetAll() ([]EventModel, error) {
	return s.eventRepo.GetAll()
}

func (s *eventService) GetById(id uuid.UUID) (*EventModel, error) {
	return s.eventRepo.GetById(id)
}

func (s *eventService) GetAllByChampionshipId(championshipId uuid.UUID) ([]EventModel, error) {
	return s.eventRepo.GetAllByChampionshipId(championshipId)
}

func (s *eventService) Create(req EventRequest) (*EventModel, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, fmt.Errorf("event name cannot be empty")
	}

	return s.eventRepo.Create(
		name,
		req.Date,
		req.ChampionshipID,
	)
}
