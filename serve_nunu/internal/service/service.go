package service

import (
	"server_go/internal/repository"
	"server_go/pkg/jwt"
	"server_go/pkg/log"
)

type Service struct {
	logger *log.Logger
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(
	tm repository.Transaction,
	logger *log.Logger,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		logger: logger,
		jwt:    jwt,
		tm:     tm,
	}
}
