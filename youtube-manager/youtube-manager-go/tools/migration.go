package main

import (
	"github.com/sirupsen/logrus"
	"youtube-manager-go/databases"
	"youtube-manager-go/models"
)

func main() {
	db, err := databases.Connect()

	sqlDb, err := db.DB()
	defer sqlDb.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Favorite{})
}
