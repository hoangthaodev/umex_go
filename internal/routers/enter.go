package routers

import (
	"umex.com/internal/routers/manage"
	"umex.com/internal/routers/user"
)

type RouterGroup struct {
	Manage manage.ManagerRouterGroup
	User   user.UserRouterGroup
}

var RouterGroupApp = new(RouterGroup)
