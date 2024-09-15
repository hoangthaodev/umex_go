package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ur *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {
	// public router
	adminRouterPublic := router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	// private router
	adminRouterPrivate := router.Group("/admin")
	// adminRouterPrivate.Use(Limiter())
	// adminRouterPrivate.Use(Authen())
	// adminRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active_user")
	}
}
