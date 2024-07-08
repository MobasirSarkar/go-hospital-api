	package utils

import (
	"log"
	"os"

	"mymodule/pkg/model"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error while loading .env", err)
	}

	dbUrl := os.Getenv("DBURL")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln("error", err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalln("error", err)
	}

	return db
}
