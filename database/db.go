package database

import (
	"github.com/alyzsa/FinPro2/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"os"
	"gorm.io/gorm"

)



var (
	host     = os.Getenv("PGHOST")
	port     = os.Getenv("PGPORT")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbname   = os.Getenv("PGDATABASE")
	db       *gorm.DB
	err      error
)


func StartDB() {
	config := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(entity.User{}, entity.Photo{}, entity.Comment{}, entity.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}



