package router

import (
	"gf-demo/app/api"
	"gf-demo/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	// 分组路由注册方式
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.CORS,
		)
		group.ALL("/hello", api.Hello)
		group.ALL("/user", api.User)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.Auth)
			group.ALL("/user/profile", api.User.Profile)
		})
	})
}
