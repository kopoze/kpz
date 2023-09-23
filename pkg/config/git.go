package config

import "log"

type GitOptions struct {
	Remote string `mapstructure:"remote"`
	Branch string `mapstructure:"branch"`
}

type GitConfig struct {
	GitOptions
}

var (
	GIT_REMOTE = "origin"
	GIT_BRANCH = "develop"
)

func NewGitConfig(opt GitOptions) GitConfig {
	config := GitConfig{
		GitOptions{
			Remote: GIT_REMOTE,
			Branch: GIT_BRANCH,
		},
	}
	newConfig, err := updateOpts(opt, config)
	if err != nil {
		log.Println(err)
	}
	config, _ = newConfig.(GitConfig)
	return config
}
