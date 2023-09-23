package app

import (
	"fmt"
	"log"

	"github.com/kopoze/kpz/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	conf := config.LoadConfig()
	var (
		host     = conf.Database.Host
		port     = conf.Database.Port
		user     = conf.Database.User
		dbname   = conf.Database.Name
		password = conf.Database.Password
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	database, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}
	err = database.AutoMigrate(&App{})

	if err != nil {
		log.Println(err)
	}

	DB = database
}
