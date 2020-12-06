package database

import (
	"github.com/stevensun369/kards/models"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"log"
)

var (
	// DBConn represents the connection to the database
	DBConn *gorm.DB
)

// InitDatabase initialises the database
func InitDatabase() {
	var err error
	dsn := "user=postgres password=StevenS369@postgres dbname=kards_dev port=5432 sslmode=disable TimeZone=Europe/Bucharest"

	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("The connection to the database failed")
	}
	log.Println("Successfully connected to the database")

	DBConn.AutoMigrate(
		&models.User{},
	)
	log.Println("AutoMigrated the models")

}