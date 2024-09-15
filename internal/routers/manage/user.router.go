package manage

import (
	"github.com/gin-gonic/gin"
	"umex.com/internal/wire"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// public router

	// this is non-dependency
	// urepo := repo.NewUserRepository()
	// uservice := service.NewUserService(urepo)
	// userHandlerNonDependency := controller.NewUserController(uservice)

	// wire
	userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := router.Group("/admin/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		// userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := router.Group("/admin/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active_user")
	}
}
