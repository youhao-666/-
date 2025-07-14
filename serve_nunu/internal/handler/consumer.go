package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	v1 "server_go/api/v1"
	"server_go/internal/model"
	"server_go/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ConsumerHandler struct {
	*Handler
	consumerService service.ConsumerService
}

func NewConsumerHandler(
	handler *Handler,
	consumerService service.ConsumerService,
) *ConsumerHandler {
	return &ConsumerHandler{
		Handler:         handler,
		consumerService: consumerService,
	}
}

func (h *ConsumerHandler) GetConsumer(ctx *gin.Context) {

}

func (h *ConsumerHandler) CreateUser(ctx *gin.Context) {
	data := ctx.PostForm("data") // 获取前端传来的data参数
	//var user model.Consumer
	var user v1.CreateUserRequest
	if err := json.Unmarshal([]byte(data), &user); err != nil { // 反序列化data参数到stu变量
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	if user.Account == "" || user.Password == "" || user.Nickname == "" {
		v1.HandleError(ctx, http.StatusBadRequest, fmt.Errorf("account, password and nickname are required"), nil)
		return
	}
	Consumer := model.Consumer{
		Account:  user.Account,
		Password: user.Password,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}
	Consumer.User_type = 1 // 这里假设用户类型为1

	err := h.consumerService.CreateConsumer(ctx, &Consumer)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *ConsumerHandler) DeleteConsumer(ctx *gin.Context) {
	user_type := GetUserTypeFromCtx(ctx)
	if user_type != 0 {
		v1.HandleError(ctx, http.StatusBadRequest, fmt.Errorf("用户类型不正确"), nil)
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // 假设 org.ID 是 uint32 类型
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	err = h.consumerService.DeleteConsumer(ctx, id)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *ConsumerHandler) ResetPassword(ctx *gin.Context) {
	user_type := GetUserTypeFromCtx(ctx)
	if user_type != 0 {
		v1.HandleError(ctx, http.StatusBadRequest, fmt.Errorf("用户类型不正确"), nil)
		return
	}
	idStr := ctx.Param("id")
	fmt.Println(user_type)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	err = h.consumerService.ResetPassword(ctx, id)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *ConsumerHandler) Login(ctx *gin.Context) {

	data := ctx.PostForm("data") // 获取前端传来的data参数

	var consumer v1.LoginRequest
	if err := json.Unmarshal([]byte(data), &consumer); err != nil { // 反序列化data参数到stu变量
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	user := model.Consumer{
		Account:  consumer.Account,
		Password: consumer.Password,
	}
	token, con, err := h.consumerService.Login(ctx, &user)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	use := v1.LoginResponse{
		ID:       con.ID,
		Account:  con.Account,
		Nickname: con.Nickname,
		Avatar:   con.Avatar,
		Token:    token,
	}

	v1.HandleSuccess(ctx, use)
}

func (h *ConsumerHandler) UpdateConsumer(ctx *gin.Context) {

	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println(id)
	data := ctx.PostForm("data") // 获取前端传来的data参数
	var user model.Consumer

	if err := json.Unmarshal([]byte(data), &user); err != nil { // 反序列化data参数到stu变量
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	// //user.ID = uint(id)
	// if user.ID != uint(id) {
	// 	v1.HandleError(ctx, http.StatusBadRequest, fmt.Errorf("id不正确"), nil)
	// 	return
	// }
	//fmt.Println(id)
	user.ID = uint(id)
	user.User_type = 1 // 这里假设用户类型为1
	err := h.consumerService.UpdateConsumer(ctx, &user)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)

}

// ==============================

func (h *ConsumerHandler) QueryAllConsumer(ctx *gin.Context) {
	user_type := GetUserTypeFromCtx(ctx)
	if user_type != 0 {
		v1.HandleError(ctx, http.StatusBadRequest, fmt.Errorf("用户类型不正确"), nil)
		return
	}

	consumers, err := h.consumerService.QueryAllConsumer(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, consumers)
}

func (h *ConsumerHandler) QueryConsumerById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	consumer, err := h.consumerService.GetConsumer(ctx, id)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	user := v1.UserResponse{
		ID:       consumer.ID,
		Account:  consumer.Account,
		Nickname: consumer.Nickname,
		Avatar:   consumer.Avatar,
	}
	v1.HandleSuccess(ctx, user)
}
