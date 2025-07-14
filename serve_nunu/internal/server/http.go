package server

import (
	"server_go/internal/handler"
	"server_go/internal/middleware"
	"server_go/pkg/jwt"
	"server_go/pkg/log"
	"server_go/pkg/server/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	consumerHandler *handler.ConsumerHandler,
	articles *handler.ArticleHandler,
	tags *handler.TagHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
	)

	s.POST("/user/login", consumerHandler.Login)
	s.POST("/user/create", consumerHandler.CreateUser)
	// 用户相关路由分组
	userGroup := s.Group("/user")
	{
		userGroup.Use(middleware.StrictAuth(jwt, logger))
		userGroup.POST("/register", userHandler.Register)
		userGroup.PUT("/update/:id", consumerHandler.UpdateConsumer)
	}

	// 管理员相关路由分组
	s.POST("/admin/login", adminHandler.Login)
	s.POST("/admin/create", adminHandler.CreateAdmin)
	adminGroup := s.Group("/admin")
	{
		adminGroup.Use(middleware.StrictAuth(jwt, logger))
		//用户管理
		adminGroup.PUT("/Reset/:id", consumerHandler.ResetPassword)
		adminGroup.GET("/user/list", consumerHandler.QueryAllConsumer)
		adminGroup.DELETE("/delete/:id", consumerHandler.DeleteConsumer)
		//标签管理
		adminGroup.POST("/tag/create", tags.CreateTag)
		adminGroup.DELETE("/tag/delete/:id", tags.DeleteTag)
		//文章管理
		adminGroup.DELETE("/article/delete/:id", articles.DeleteArticle)
	}
	s.POST("/article/create", articles.CreateArticle)
	s.GET("/tag/list", tags.QueryAllTag)
	s.GET("/user/:id", consumerHandler.QueryConsumerById)

	article := s.Group("/article")
	{
		//文章查询
		//article.Use(middleware.StrictAuth(jwt, logger))
		article.GET("/listByTag/:tag_id", articles.QueryArticleByTag)
		article.GET("/quaryall", articles.QueryAllArticle)
		article.GET("/listByUser/:user_id", articles.QueryArticleByUserID)
	}

	return s
}
