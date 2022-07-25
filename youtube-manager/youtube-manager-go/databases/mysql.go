package database

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Connect() (db *gorm.DB, err error) {
	err = godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	db, err = gorm.Open(mysql.Open(
		os.Getenv("DB_USERNAME")+
			":"+
			os.Getenv("DB_PASSWORD")+
			"@tcp("+
			os.Getenv("DB_HOST")+
			":"+
			os.Getenv("DB_PORT")+
			")/"+
			os.Getenv("DB_DATABASE")+
			"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		logrus.Fatal(err)
	}

	return db, err
}
