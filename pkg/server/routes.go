package server

import (
	"github.com/kopoze/kpz/pkg/reverseproxy"
)

// import proxy github.com/kopoze/kopoze-tools/pkg/libreverseproxy"

func RegisterRoutes(r *Routes) {
	r.router.Use(reverseproxy.ReverseProxy())

	cli := r.router.Group("/cli")

	appRoute := cli.Group("/apps")
	appRoute.GET("/", ListApps)
	appRoute.POST("/", CreateApp)
	appRoute.GET("/:id", FindApp)
	appRoute.PATCH("/:id", UpdateApp)
	appRoute.DELETE("/:id", DeleteApp)
}
