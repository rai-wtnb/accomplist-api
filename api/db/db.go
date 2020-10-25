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

// Init is makes connection to psql.
func Init() {
	db, err = gorm.Open("postgres", "host=db port=5432 user=accomplist dbname=accomplist password=accomplist-password sslmode=disable")
	if err != nil {
		panic(err)
	}

	// todo
	autoMigration()
	// user := models.User{
	// 	ID:          2,
	// 	Name:        "bbb",
	// 	Email:       "bbb@bbb.com",
	// 	Password:    "bbbbbb",
	// 	Twitter:     "mmuu_kkuu",
	// 	Description: "aiueokakikukeko",
	// 	ImgPath:     "https://via.placeholder.com/150",
	// }
	// db.Create(&user)
}

// GetDB is gets db.
func GetDB() *gorm.DB {
	return db
}

// Close is closed db.
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// autoMigration is migrates in accordance with models
func autoMigration() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.List{})
	db.AutoMigrate(&models.Like{})
}
