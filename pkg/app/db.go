package app

import (
	"fmt"
	"log"

	"github.com/kopoze/kpz/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// TODO: Get value from config
	var (
		host     = utils.GetEnv("DB_HOST")
		port     = utils.GetEnv("DB_PORT")
		user     = utils.GetEnv("DB_USER")
		dbname   = utils.GetEnv("DB_NAME")
		password = utils.GetEnv("DB_PASSWORD")
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
