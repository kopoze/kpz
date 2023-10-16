package config

import "log"

type DockerConfig struct {
	DockerOptions
}

type DockerOptions struct {
	Shell string `mapstructure:"shell"`
}

var (
	DOCKER_SHELL = "bash"
)

// Create new DockerConfig.
func NewDockerConfig(opt DockerOptions) DockerConfig {
	config := DockerConfig{
		DockerOptions{
			Shell: DOCKER_SHELL,
		},
	}
	newConfig, err := updateOpts(opt, config)
	if err != nil {
		log.Println(err)
	}
	config, _ = newConfig.(DockerConfig)
	return config
}
