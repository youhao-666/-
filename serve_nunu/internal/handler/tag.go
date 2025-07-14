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

type TagHandler struct {
	*Handler
	tagService service.TagService
}

func NewTagHandler(
	handler *Handler,
	tagService service.TagService,
) *TagHandler {
	return &TagHandler{
		Handler:    handler,
		tagService: tagService,
	}
}

func (h *TagHandler) CreateTag(ctx *gin.Context) {
	user_type := GetUserTypeFromCtx(ctx)
	if user_type != 0 {
		v1.HandleError(ctx, http.StatusBadRequest, fmt.Errorf("用户类型不正确"), nil)
		return
	}
	data := ctx.PostForm("data") // 获取前端传来的data参数
	var tag model.Tag
	if err := json.Unmarshal([]byte(data), &tag); err != nil { // 反序列化data参数到stu变量
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	if err := h.tagService.CreateTag(ctx, &tag); err != nil { // 调用service层的GetTag方法
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *TagHandler) QueryAllTag(ctx *gin.Context) {
	tags, err := h.tagService.QueryAllTag(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	var tagList []v1.Tag
	for _, tag := range tags {
		tagList = append(tagList, v1.Tag{ID: int(tag.ID), TagName: tag.TagName})
	}
	v1.HandleSuccess(ctx, tagList)
}

func (h *TagHandler) DeleteTag(ctx *gin.Context) {
	user_type := GetUserTypeFromCtx(ctx)
	if user_type != 0 {
		v1.HandleError(ctx, http.StatusBadRequest, fmt.Errorf("用户类型不正确"), nil)
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	userID := uint(id)
	if err := h.tagService.DeleteTag(ctx, userID); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)

}
