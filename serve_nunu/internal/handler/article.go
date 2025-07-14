package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	v1 "server_go/api/v1"
	"server_go/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	*Handler
	articleService service.ArticleService
}

func NewArticleHandler(
	handler *Handler,
	articleService service.ArticleService,
) *ArticleHandler {
	return &ArticleHandler{
		Handler:        handler,
		articleService: articleService,
	}
}

func (h *ArticleHandler) GetArticle(ctx *gin.Context) {

}

func (h *ArticleHandler) CreateArticle(ctx *gin.Context) {
	data := ctx.PostForm("data") // 获取前端传来的data参数
	var articleWithTags v1.ArticleWithTags
	if err := json.Unmarshal([]byte(data), &articleWithTags); err != nil { // 反序列化data参数到stu变量
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	if articleWithTags.Title == "" || articleWithTags.Content == "" || articleWithTags.UserId <= 0 {
		v1.HandleError(ctx, http.StatusBadRequest, fmt.Errorf("title, content, userId不能为空"), nil)
		return
	}
	if err := h.articleService.CreateArticle(ctx, &articleWithTags); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	} // 调用service层的CreateArticle方法
	v1.HandleSuccess(ctx, nil)
}

func (h *ArticleHandler) QueryArticleByTag(ctx *gin.Context) {
	// user_type := GetUserIdFromCtx(ctx)
	// user_id := GetUserIdFromCtx(ctx)
	idStr := ctx.Param("tag_id")
	id, err := strconv.Atoi(idStr)
	tagID := uint(id)

	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	articles, err := h.articleService.QueryArticleByTag(ctx, tagID)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, articles)

}

func (h *ArticleHandler) QueryAllArticle(ctx *gin.Context) {
	articles, err := h.articleService.QueryAllArticle(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, articles)
}

func (h *ArticleHandler) QueryArticleByUserID(ctx *gin.Context) {
	idStr := ctx.Param("user_id")
	id, err := strconv.Atoi(idStr)
	userID := uint(id)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	//fmt.Println(userID)
	articles, err := h.articleService.QueryArticleByUserID(ctx, userID)
	fmt.Println(articles)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, articles)
}

func (h *ArticleHandler) DeleteArticle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	ArticleID := uint(id)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	if err := h.articleService.DeleteArticle(ctx, ArticleID); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}
