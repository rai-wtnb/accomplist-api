package db

import (
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
	db, err = gorm.Open("postgres", "host=db port=5432 user=accomplist dbname=accomplist password=accomplist-password sslmode=disable")
	if err != nil {
		panic(err)
	}

	// todo
	autoMigration()
	user := models.User{
		Name:     "aaa",
		Email:    "aaa@aaa.com",
		Password: "aaaaaa",
	}
	db.Create(&user)
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
