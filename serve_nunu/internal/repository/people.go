package repository

import (
    "context"
	"server_go/internal/model"
)

type PeopleRepository interface {
	GetPeople(ctx context.Context, id int64) (*model.People, error)
}

func NewPeopleRepository(
	repository *Repository,
) PeopleRepository {
	return &peopleRepository{
		Repository: repository,
	}
}

type peopleRepository struct {
	*Repository
}

func (r *peopleRepository) GetPeople(ctx context.Context, id int64) (*model.People, error) {
	var people model.People

	return &people, nil
}
