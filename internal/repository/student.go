package repository

import (
    "context"
	"/internal/model"
)

type StudentRepository interface {
	GetStudent(ctx context.Context, id int64) (*model.Student, error)
}

func NewStudentRepository(
	repository *Repository,
) StudentRepository {
	return &studentRepository{
		Repository: repository,
	}
}

type studentRepository struct {
	*Repository
}

func (r *studentRepository) GetStudent(ctx context.Context, id int64) (*model.Student, error) {
	var student model.Student

	return &student, nil
}
