package client

import (
	"fmt"

	"github.com/kopoze/kpz/pkg/config"
)

func BuildURL() string {
	conf := config.LoadConfig()

	var url string
	switch conf.Kopoze.Mode {
	case "remote":
		url = conf.Kopoze.Host

	default: // default mode is 'local'
		url = fmt.Sprintf("http://localhost:%s/cli/apps/", conf.Kopoze.Port)
	}

	return url
}
