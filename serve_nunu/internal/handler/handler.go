package handler

import (
	"server_go/pkg/jwt"
	"server_go/pkg/log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(
	logger *log.Logger,
) *Handler {
	return &Handler{
		logger: logger,
	}
}

// func GetUserIdFromCtx(ctx *gin.Context) uint {
// 	v, exists := ctx.Get("claims")
// 	if !exists {
// 		return 0
// 	}
// 	return v.(*jwt.MyCustomClaims).UserId
// }

func GetUserIdFromCtx(ctx *gin.Context) uint {
	v, exists := ctx.Get("claims")
	if !exists {
		return 0
	}
	claims, ok := v.(*jwt.MyCustomClaims)
	if !ok {
		return 0
	}
	return claims.UserId
}

func GetUserTypeFromCtx(ctx *gin.Context) int {
	v, exists := ctx.Get("claims")
	if !exists {
		return 0
	}
	claims, ok := v.(*jwt.MyCustomClaims)
	if !ok {
		return -1
	}
	return claims.User_type
}
