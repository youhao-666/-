package handler

import (
	"fmt"
	"net/http"
	v1 "server_go/api/v1"
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

	var tag v1.TagCreate
	if err := ctx.ShouldBindJSON(&tag); err != nil {
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
