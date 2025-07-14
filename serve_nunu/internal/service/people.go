package service

import (
    "context"
	"server_go/internal/model"
	"server_go/internal/repository"
)

type PeopleService interface {
	GetPeople(ctx context.Context, id int64) (*model.People, error)
}
func NewPeopleService(
    service *Service,
    peopleRepository repository.PeopleRepository,
) PeopleService {
	return &peopleService{
		Service:        service,
		peopleRepository: peopleRepository,
	}
}

type peopleService struct {
	*Service
	peopleRepository repository.PeopleRepository
}

func (s *peopleService) GetPeople(ctx context.Context, id int64) (*model.People, error) {
	return s.peopleRepository.GetPeople(ctx, id)
}
