package sports

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type SportService interface {
	GetAll() ([]SportModel, error)
	GetById(id uuid.UUID) (*SportModel, error)
	Create(req CreateSportRequest) (*SportModel, error)
}

type sportService struct {
	sportRepo SportRepository
}

func NewSportService(sportRepo SportRepository) SportService {
	return &sportService{sportRepo: sportRepo}
}

func (s *sportService) GetAll() ([]SportModel, error) {
	return s.sportRepo.GetAll()
}

func (s *sportService) GetById(id uuid.UUID) (*SportModel, error) {
	return s.sportRepo.GetById(id)
}

func (s *sportService) Create(req CreateSportRequest) (*SportModel, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, fmt.Errorf("Sport name cannot be empty")
	}

	return s.sportRepo.Create(name)
}
