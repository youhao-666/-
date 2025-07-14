package service

import (
	"context"
	v1 "server_go/api/v1"
	"server_go/internal/model"
	"server_go/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *v1.RegisterRequest) error
}

func NewUserService(
	service *Service,
	userRepo repository.UserRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest) error {
	user, err := s.userRepo.GetByAccount(ctx, req.Account)
	if err != nil {
		return v1.ErrInternalServerError
	}
	if user != nil {
		return v1.ErrAccountAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user = &model.User{
		Account:  req.Account,
		Password: string(hashedPassword),
	}
	if err = s.userRepo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}
