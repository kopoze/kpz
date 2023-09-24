package config

import "log"

type KopozeOptions struct {
	Version string `mapstructure:"version"`
	Domain  string `mapstructure:"domain"`
	Port    string `mapstructure:"port"`
	Mode    string `mapstructure:"mode"`
}

type KopozeConfig struct {
	KopozeOptions
}

var (
	KPZ_PORT   = "8080"
	KPZ_DOMAIN = "project.mg"
	KPZ_MODE   = "local"
)

func NewKopozeConfig(opt KopozeOptions) KopozeConfig {
	config := KopozeConfig{
		KopozeOptions{
			Version: APP_VERSION,
			Domain:  KPZ_DOMAIN,
			Port:    KPZ_PORT,
			Mode:    KPZ_MODE,
		},
	}
	newConfig, err := updateOpts(opt, config)
	if err != nil {
		log.Println(err)
	}
	config, _ = newConfig.(KopozeConfig)
	return config
}
