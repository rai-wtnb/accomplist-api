package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rai-wtnb/accomplist-api/models"
)

var (
	db  *gorm.DB
	err error
)

// Init makes connection to psql.
func Init() {
	connection := fmt.Sprintf(
		"host=%s port=5432 user=accomplist dbname=accomplist password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PASS"),
	)
	db, err = gorm.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	autoMigration()
}

// GetDB gets db.
func GetDB() *gorm.DB {
	return db
}

// Close closed db.
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// autoMigration migrates in accordance with models
func autoMigration() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.List{})
	db.AutoMigrate(&models.Like{})
	db.AutoMigrate(&models.Feedback{})
}
