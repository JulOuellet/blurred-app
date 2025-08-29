package seasons

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type SeasonService interface {
	GetAll() ([]SeasonModel, error)
	GetById(id uuid.UUID) (*SeasonModel, error)
	Create(req SeasonRequest) (*SeasonModel, error)
}

type seasonService struct {
	seasonRepo SeasonRepository
}

func NewSeasonService(seasonRepo SeasonRepository) SeasonService {
	return &seasonService{seasonRepo: seasonRepo}
}

func (s *seasonService) GetAll() ([]SeasonModel, error) {
	return s.seasonRepo.GetAll()
}

func (s *seasonService) GetById(id uuid.UUID) (*SeasonModel, error) {
	return s.seasonRepo.GetById(id)
}

func (s *seasonService) Create(req SeasonRequest) (*SeasonModel, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, fmt.Errorf("Season name cannot be empty")
	}

	if req.EndDate.Before(req.StartDate) {
		return nil, fmt.Errorf("end date cannot be before start date")
	}

	return s.seasonRepo.Create(name, req.StartDate, req.EndDate, req.SportID)
}
