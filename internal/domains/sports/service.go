package sports

import (
	"fmt"
	"strings"

	"github.com/JulOuellet/sportlight/internal/domains/seasons"
	"github.com/google/uuid"
)

type SportService interface {
	GetAll() ([]SportModel, error)
	GetById(id uuid.UUID) (*SportModel, error)
	Create(req SportRequest) (*SportModel, error)
	Update(id uuid.UUID, req SportRequest) (*SportModel, error)
	GetSportWithSeasons(id uuid.UUID) (*SportWithSeasons, error)
}

type sportService struct {
	sportRepo  SportRepository
	seasonrepo seasons.SeasonRepository
}

func NewSportService(sportRepo SportRepository, seasonRepo seasons.SeasonRepository) SportService {
	return &sportService{
		sportRepo:  sportRepo,
		seasonrepo: seasonRepo,
	}
}

func (s *sportService) GetAll() ([]SportModel, error) {
	return s.sportRepo.GetAll()
}

func (s *sportService) GetById(id uuid.UUID) (*SportModel, error) {
	return s.sportRepo.GetById(id)
}

func (s *sportService) Create(req SportRequest) (*SportModel, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, fmt.Errorf("Sport name cannot be empty")
	}

	return s.sportRepo.Create(name)
}

func (s *sportService) Update(id uuid.UUID, req SportRequest) (*SportModel, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, fmt.Errorf("Sport name cannot be empty")
	}

	return s.sportRepo.Update(id, name)
}

func (s *sportService) GetSportWithSeasons(id uuid.UUID) (*SportWithSeasons, error) {
	sport, err := s.sportRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	seasons, err := s.seasonrepo.GetAllBySportId(id)
	if err != nil {
		return nil, err
	}

	return &SportWithSeasons{
		SportModel: *sport,
		Seasons:    seasons,
	}, nil
}
