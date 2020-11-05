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
	// user := models.User{
	// 	ID:          "aaaaaa",
	// 	Name:        "aaa",
	// 	Email:       "aaa@aaa.com",
	// 	Password:    "aaaaaa",
	// 	Twitter:     "mmuu_kkuu",
	// 	Description: "エンジニア志望の22卒",
	// 	Img:         "https://via.placeholder.com/60",
	// 	Lists: []models.List{{
	// 		Content: "海外旅行をする",
	// 		Done:    true,
	// 		Feedback: models.Feedback{
	// 			ImgPath: "https://via.placeholder.com/60",
	// 			Title:   "カナダへの旅行!",
	// 			Body:    "念願の海外旅行を達成しました。Vlogも撮ってきた!",
	// 		},
	// 	}},
	// }
	// db.Create(&user)
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
