package models

import (
	"encoding/json"
	"log"
	"time"

	database "github.com/MrHenri/gonews/db"
)

type News struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"size:255;not null;unique" json:"content"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (n *News) GetAllNews() ([]News, error) {
	connection, _ := database.Connect()
	defer connection.Close()

	newsList := []News{}
	reply, err := database.Get("news")

	if err != nil {
		log.Println("Searching on mySql")
		err := connection.Debug().Model(&News{}).Limit(100).Find(&newsList).Error
		if err != nil {
			return []News{}, err
		}
		newsBytes, _ := json.Marshal(newsList)
		database.Set("news", newsBytes)
		return newsList, nil
	}

	log.Println("Searching on Redis")
	json.Unmarshal(reply, &newsList)
	return newsList, nil
}

func (n *News) AddNews() (*News, error) {
	connection, _ := database.Connect()
	defer connection.Close()

	err := connection.Debug().Model(&News{}).Create(&n).Error
	if err != nil {
		return &News{}, err
	}
	database.Flush("news")
	return n, nil
}
