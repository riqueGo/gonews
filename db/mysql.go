package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DB_NAME = "gonews"
	DB_HOST = "db"
	DB_USER = "gonewsuser"
	DB_PASS = "gonewsps"
	DB_PORT = 3306
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME))

	if err != nil {
		fmt.Errorf("Error connecting")
	}

	return db, err
}
