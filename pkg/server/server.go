package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kopoze/kpz/pkg/app"
	"github.com/kopoze/kpz/pkg/config"
)

type Routes struct {
	router *gin.Engine
}

func Serve() {
	conf := config.LoadConfig()

	r := Routes{router: gin.Default()}

	app.ConnectDB()

	RegisterRoutes(&r)

	r.router.Run(fmt.Sprintf(":%s", conf.Kopoze.Port))
}
