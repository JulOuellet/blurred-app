package seasons

type SeasonSerivce interface {
	GetAll() ([]SeasonModel, error)
}

type seasonService struct {
	repository SeasonRepository
}

func NewSeasonService(repository SeasonRepository) SeasonSerivce {
	return &seasonService{repository: repository}
}

func (s *seasonService) GetAll() ([]SeasonModel, error) {
	seasons, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return seasons, nil
}
