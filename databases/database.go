package databases

import (
	"github/asuzukosi/fiber-api/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	log.Println("Successfully connected to the database")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations....")
	db.AutoMigrate(&models.User{}, &models.Order{}, &models.Product{})

	Database.Db = db
}
