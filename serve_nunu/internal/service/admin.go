package service

import (
	"context"
	"errors"
	v1 "server_go/api/v1"
	"server_go/internal/model"
	"server_go/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	GetAdmin(ctx context.Context, id int64) (*model.Admin, error)
	CreateAdmin(ctx context.Context, admin *model.Admin) error

	Login(ctx context.Context, admin *model.Admin) (string, *model.Admin, error)
}

func NewAdminService(
	service *Service,
	adminRepository repository.AdminRepository,
) AdminService {
	return &adminService{
		Service:         service,
		adminRepository: adminRepository,
	}
}

type adminService struct {
	*Service
	adminRepository repository.AdminRepository
}

func (s *adminService) GetAdmin(ctx context.Context, id int64) (*model.Admin, error) {
	return s.adminRepository.GetAdmin(ctx, id)
}

func (s *adminService) CreateAdmin(ctx context.Context, admin *model.Admin) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin.Password = string(hashedPassword)

	err = s.adminRepository.CreateAdmin(ctx, admin)
	if err != nil {
		return err
	}
	return nil
}

func (s *adminService) Login(ctx context.Context, admin *model.Admin) (string, *model.Admin, error) {
	account := admin.Account
	pass, err := s.adminRepository.CheckAccount(ctx, account)
	if err != nil {
		return "", nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(pass.Password), []byte(admin.Password))
	if err != nil {
		return "", nil, errors.New("password error")
	}

	//2.生成token
	token, err := s.jwt.GenToken(admin.ID, time.Now().Add(time.Hour*24*7), admin.User_type)
	if err != nil {
		return "", nil, v1.ErrInternalServerError
	}

	//3.返回token
	return token, pass, nil

}
