package migration

import (
	"log"

	db "github.com/MrHenri/gonews/db"
	"github.com/MrHenri/gonews/models"
)

func AutoMigration() bool {
	db, err := db.Connect()
	defer db.Close()

	if err != nil {
		log.Println("It was not possible migrate, database is not available")
		return false
	}
	db.Debug().AutoMigrate(&models.News{})
	//Auto create table based on Model
	return true
}
