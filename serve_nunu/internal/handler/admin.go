package handler

import (
	"encoding/json"
	"net/http"
	v1 "server_go/api/v1"
	"server_go/internal/model"
	"server_go/internal/service"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	*Handler
	adminService service.AdminService
}

func NewAdminHandler(
	handler *Handler,
	adminService service.AdminService,
) *AdminHandler {
	return &AdminHandler{
		Handler:      handler,
		adminService: adminService,
	}
}

func (h *AdminHandler) CreateAdmin(ctx *gin.Context) {
	data := ctx.PostForm("data") // 获取前端传来的data参数
	var user model.Admin
	if err := json.Unmarshal([]byte(data), &user); err != nil { // 反序列化data参数到stu变量
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	user.User_type = 0 // 这里假设用户类型为1

	err := h.adminService.CreateAdmin(ctx, &user)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *AdminHandler) Login(ctx *gin.Context) {
	data := ctx.PostForm("data") // 获取前端传来的data参数

	var userr v1.LoginRequest
	if err := json.Unmarshal([]byte(data), &userr); err != nil { // 反序列化data参数到stu变量
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	user := model.Admin{
		Account:  userr.Account,
		Password: userr.Password,
	}
	token, admin, err := h.adminService.Login(ctx, &user)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	adminRes := v1.LoginResponse{
		ID:       admin.ID,
		Account:  admin.Account,
		Nickname: admin.Nickname,
		Avatar:   admin.Avatar,
		Token:    token,
	}
	v1.HandleSuccess(ctx, adminRes)
}
