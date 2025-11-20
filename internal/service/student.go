package service

import (
    "context"
	"/internal/model"
	"/internal/repository"
)

type StudentService interface {
	GetStudent(ctx context.Context, id int64) (*model.Student, error)
}
func NewStudentService(
    service *Service,
    studentRepository repository.StudentRepository,
) StudentService {
	return &studentService{
		Service:        service,
		studentRepository: studentRepository,
	}
}

type studentService struct {
	*Service
	studentRepository repository.StudentRepository
}

func (s *studentService) GetStudent(ctx context.Context, id int64) (*model.Student, error) {
	return s.studentRepository.GetStudent(ctx, id)
}
