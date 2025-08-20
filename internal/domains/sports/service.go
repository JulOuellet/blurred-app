package sports

type SportService interface {
	GetAll() ([]SportModel, error)
}

type sportService struct {
	sportRepo SportRepository
}

func NewSportService(sportRepo SportRepository) SportService {
	return &sportService{sportRepo: sportRepo}
}

func (s *sportService) GetAll() ([]SportModel, error) {
	sports, err := s.sportRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return sports, nil
}
