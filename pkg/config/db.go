package config

import "log"

type DBOptions struct {
	Engine   string `mapstructure:"engine"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type DBConfig struct {
	DBOptions //`mapstructure:",squash"`

}

var (
	DB_ENGINE   = "postgresql"
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "root"
	DB_PASSWORD = "root"
)

func NewDBConfig(opt DBOptions) DBConfig {
	config := DBConfig{
		DBOptions{
			Engine:   DB_ENGINE,
			Host:     DB_HOST,
			Port:     DB_PORT,
			User:     DB_USER,
			Password: DB_PASSWORD,
		},
	}

	newConfig, err := updateOpts(opt, config)
	if err != nil {
		log.Println(err)
	}
	config, _ = newConfig.(DBConfig)
	return config
}
