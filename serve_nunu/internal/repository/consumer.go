package repository

import (
	"context"
	"errors"
	v1 "server_go/api/v1"
	"server_go/internal/model"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type ConsumerRepository interface {
	GetConsumer(ctx context.Context, id int64) (*model.Consumer, error)
	CreateConsumer(ctx context.Context, consumer *model.Consumer) error
	DeleteConsumer(ctx context.Context, id uint64) error
	QueryConsumerById(ctx context.Context, id uint64) (*model.Consumer, error)
	ResetPassword(ctx context.Context, id uint64, consumer *model.Consumer) error

	CheckAccount(ctx context.Context, account string) (*model.Consumer, error)
	UpdateConsumer(ctx context.Context, consumer *model.Consumer) error
	QueryAllConsumer(ctx context.Context) ([]*model.Consumer, error)
}

func NewConsumerRepository(
	repository *Repository,
) ConsumerRepository {
	return &consumerRepository{
		Repository: repository,
	}
}

type consumerRepository struct {
	*Repository
}

func (r *consumerRepository) GetConsumer(ctx context.Context, id int64) (*model.Consumer, error) {
	var consumer model.Consumer
	if err := r.db.Table("user").Where("id = ?", id).First(&consumer).Error; err != nil {
		return nil, err
	}
	return &consumer, nil
}

func (r *consumerRepository) CreateConsumer(ctx context.Context, consumer *model.Consumer) error {
	if err := r.DB(ctx).Table("user").Create(consumer).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return v1.ErrUserAlreadyUse
			}
			return err
		}
	}
	return nil
}

func (r *consumerRepository) DeleteConsumer(ctx context.Context, id uint64) error {
	// 方式1：直接删除（推荐）
	result := r.db.Table("user").Where("id = ?", id).Delete(&model.Consumer{})
	if result.Error != nil {
		return result.Error
	}

	// 检查是否确实删除了记录
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *consumerRepository) QueryConsumerById(ctx context.Context, id uint64) (*model.Consumer, error) {
	var consumer model.Consumer
	result := r.db.Table("user").Where("id = ?", id).First(&consumer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, v1.ErrIdNotFound
		}
		return nil, result.Error
	}
	return &consumer, nil
}

func (r *consumerRepository) ResetPassword(ctx context.Context, id uint64, consumer *model.Consumer) error {
	if err := r.db.Table("user").Where("id = ?", id).Updates(consumer).Error; err != nil {
		return err
	}
	return nil

}

func (r *consumerRepository) CheckAccount(ctx context.Context, account string) (*model.Consumer, error) {
	var consumer model.Consumer
	if err := r.db.Table("user").Where("account = ? AND user_type = 1", account).First(&consumer).Error; err != nil {
		return nil, err
	}
	return &consumer, nil
}

func (r *consumerRepository) UpdateConsumer(ctx context.Context, consumer *model.Consumer) error {
	if err := r.db.Table("user").Where("id = ?", consumer.ID).Updates(consumer).Error; err != nil {
		return err
	}
	return nil
}

func (r *consumerRepository) QueryAllConsumer(ctx context.Context) ([]*model.Consumer, error) {
	var consumers []*model.Consumer
	if err := r.db.Table("user").Where("user_type = ?", 1).Find(&consumers).Error; err != nil {
		return nil, err
	}
	return consumers, nil
}
