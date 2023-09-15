package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kopoze/kpz/pkg/app"
)

type Routes struct {
	router *gin.Engine
}

func Serve() {
	r := Routes{router: gin.Default()}

	app.ConnectDB()

	RegisterRoutes(&r)

	r.router.Run(":8080")
}
