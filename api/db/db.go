package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rai-wtnb/accomplist-api/models"
)

var (
	Db  *gorm.DB
	err error
)

// Init makes connection to psql.
func Init() {
	conn := fmt.Sprintf(
		"host=%s port=5432 user=accomplist dbname=accomplist password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PASS"),
	)
	Db, err = gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	autoMigration()
}

// GetDB gets db.
func GetDB() *gorm.DB {
	return Db
}

// Close closed db.
func Close() {
	if err := Db.Close(); err != nil {
		panic(err)
	}
}

// autoMigration migrates in accordance with models
func autoMigration() {
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.List{})
	Db.AutoMigrate(&models.Like{})
	Db.AutoMigrate(&models.Feedback{})
}
