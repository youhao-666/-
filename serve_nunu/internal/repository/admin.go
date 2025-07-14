package repository

import (
	"context"
	"server_go/internal/model"
)

type AdminRepository interface {
	GetAdmin(ctx context.Context, id int64) (*model.Admin, error)
	CreateAdmin(ctx context.Context, admin *model.Admin) error
	CheckAccount(ctx context.Context, account string) (*model.Admin, error)
}

func NewAdminRepository(
	repository *Repository,
) AdminRepository {
	return &adminRepository{
		Repository: repository,
	}
}

type adminRepository struct {
	*Repository
}

func (r *adminRepository) GetAdmin(ctx context.Context, id int64) (*model.Admin, error) {
	var admin model.Admin

	return &admin, nil
}

func (r *adminRepository) CreateAdmin(ctx context.Context, admin *model.Admin) error {
	if err := r.db.Table("user").Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) CheckAccount(ctx context.Context, account string) (*model.Admin, error) {
	var admin model.Admin
	if err := r.db.Table("user").Where("account = ? AND user_type = 0", account).First(&admin).Error; err != nil {
		return nil, err
	}
	// if err := r.db.Where("account = ?", account).First(&admin).Error; err != nil {
	//     return "", err
	// }
	// if err := r.db.Table("admin").Where("account = ?", account).Order("").First(&admin).Error; err != nil {
	//     return "", err
	// }
	return &admin, nil
}
