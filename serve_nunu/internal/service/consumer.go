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

type ConsumerService interface {
	GetConsumer(ctx context.Context, id int64) (*model.Consumer, error)
	CreateConsumer(ctx context.Context, consumer *model.Consumer) error
	DeleteConsumer(ctx context.Context, id uint64) error
	ResetPassword(ctx context.Context, id uint64) error

	Login(ctx context.Context, consumer *model.Consumer) (string, *model.Consumer, error)
	UpdateConsumer(ctx context.Context, consumer *model.Consumer) error
	QueryAllConsumer(ctx context.Context) ([]*model.Consumer, error)
}

func NewConsumerService(
	service *Service,
	consumerRepository repository.ConsumerRepository,
) ConsumerService {
	return &consumerService{
		Service:            service,
		consumerRepository: consumerRepository,
	}
}

type consumerService struct {
	*Service
	consumerRepository repository.ConsumerRepository
}

func (s *consumerService) GetConsumer(ctx context.Context, id int64) (*model.Consumer, error) {
	return s.consumerRepository.GetConsumer(ctx, id)
}

func (s *consumerService) CreateConsumer(ctx context.Context, consumer *model.Consumer) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(consumer.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	consumer.Password = string(hashedPassword)

	if err := s.consumerRepository.CreateConsumer(ctx, consumer); err != nil {
		return err
	}
	return nil
}

func (s *consumerService) DeleteConsumer(ctx context.Context, id uint64) error {
	if _, err := s.consumerRepository.QueryConsumerById(ctx, id); err != nil {
		return v1.ErrIdNotFound
	}
	if err := s.consumerRepository.DeleteConsumer(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *consumerService) ResetPassword(ctx context.Context, id uint64) error {
	var consumer *model.Consumer
	var err error
	if consumer, err = s.consumerRepository.QueryConsumerById(ctx, id); err != nil {
		return v1.ErrIdNotFound
	}
	Password := "123456"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	consumer.Password = string(hashedPassword)
	if err := s.consumerRepository.ResetPassword(ctx, id, consumer); err != nil {
		return err
	}
	return nil

}

func (s *consumerService) Login(ctx context.Context, consumer *model.Consumer) (string, *model.Consumer, error) {
	account := consumer.Account
	con, err := s.consumerRepository.CheckAccount(ctx, account)
	if err != nil {
		return "", nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(con.Password), []byte(consumer.Password))
	if err != nil {
		return "", nil, errors.New("密码错误")
	}

	//2.生成token
	//fmt.Println(con.ID)
	token, err := s.jwt.GenToken(con.ID, time.Now().Add(time.Hour*24*7), consumer.User_type)
	if err != nil {
		return "", nil, v1.ErrInternalServerError
	}

	//3.返回token
	return token, con, nil
}

func (s *consumerService) UpdateConsumer(ctx context.Context, consumer *model.Consumer) error {
	id := uint64(consumer.ID)
	if _, err := s.consumerRepository.QueryConsumerById(ctx, id); err != nil {
		return v1.ErrIdNotFound
	}
	if consumer.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(consumer.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		consumer.Password = string(hashedPassword)
	}

	if err := s.consumerRepository.UpdateConsumer(ctx, consumer); err != nil {
		return err
	}
	return nil
}

func (s *consumerService) QueryAllConsumer(ctx context.Context) ([]*model.Consumer, error) {
	return s.consumerRepository.QueryAllConsumer(ctx)
}
