package database

import (
	"fmt"
	"jwt-go/app/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "book"
	db       *gorm.DB
	err      error
)

func Connect() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(entity.Users{}, entity.Product{})
	return db, nil

}
