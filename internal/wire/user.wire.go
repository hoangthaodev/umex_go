//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"umex.com/internal/controller"
	"umex.com/internal/repo"
	"umex.com/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
