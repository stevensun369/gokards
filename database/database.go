package database

import (
	"github.com/stevensun369/kards/models"
	"github.com/stevensun369/kards/conf"

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

	DBConn, err = gorm.Open(postgres.Open(conf.DSN), &gorm.Config{})

	if err != nil {
		log.Println("The connection to the database failed")
	}
	log.Println("Successfully connected to the database")

	DBConn.AutoMigrate(
		&models.User{},
		&models.Kard{},
	)
		
	log.Println("AutoMigrated the models")
}