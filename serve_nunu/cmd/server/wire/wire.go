//go:build wireinject
// +build wireinject

package wire

import (
	"server_go/internal/handler"
	"server_go/internal/repository"
	"server_go/internal/server"
	"server_go/internal/service"
	"server_go/pkg/app"
	"server_go/pkg/jwt"
	"server_go/pkg/log"
	"server_go/pkg/server/http"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewAdminRepository,
	repository.NewConsumerRepository,
	repository.NewTagRepository,
	repository.NewArticleRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewAdminService,
	service.NewConsumerService,
	service.NewTagService,
	service.NewArticleService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewAdminHandler,
	handler.NewConsumerHandler,
	handler.NewTagHandler,
	handler.NewArticleHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
)

// build App
func newApp(
	httpServer *http.Server,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		jwt.NewJwt,
		newApp,
	))
}
