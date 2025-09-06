package championships

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type ChampionshipService interface {
	GetAll() ([]ChampionshipModel, error)
	GetById(id uuid.UUID) (*ChampionshipModel, error)
	Create(req ChampionshipRequest) (*ChampionshipModel, error)
	GetAllBySeasonId(seasonId uuid.UUID) ([]ChampionshipModel, error)
}

type championshipService struct {
	championshipRepo ChampionshipRepository
}

func NewChampionshipService(championshipRepo ChampionshipRepository) ChampionshipService {
	return &championshipService{championshipRepo: championshipRepo}
}

func (s *championshipService) GetAll() ([]ChampionshipModel, error) {
	return s.championshipRepo.GetAll()
}

func (s *championshipService) GetById(id uuid.UUID) (*ChampionshipModel, error) {
	return s.championshipRepo.GetById(id)
}

func (s *championshipService) GetAllBySeasonId(seasonId uuid.UUID) ([]ChampionshipModel, error) {
	return s.championshipRepo.GetAllBySeasonId(seasonId)
}

func (s *championshipService) Create(req ChampionshipRequest) (*ChampionshipModel, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, fmt.Errorf("championship name cannot be empty")
	}

	referenceImgURL := strings.TrimSpace(req.ReferenceImgURL)
	if referenceImgURL == "" {
		return nil, fmt.Errorf("reference image URL cannot be empty")
	}

	if req.StartDate != nil && req.EndDate != nil {
		if req.EndDate.Before(*req.StartDate) {
			return nil, fmt.Errorf("end date cannot be before start date")
		}
	}

	var organization string
	if req.Organization != nil {
		organization = strings.TrimSpace(*req.Organization)
	}

	var description string
	if req.Description != nil {
		description = strings.TrimSpace(*req.Description)
	}

	return s.championshipRepo.Create(
		name,
		organization,
		req.StartDate,
		req.EndDate,
		req.SeasonID,
		description,
		referenceImgURL,
	)
}
