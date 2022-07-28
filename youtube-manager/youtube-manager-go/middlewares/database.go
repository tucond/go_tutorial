package middlewares

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"youtube-manager-go/databases"
)

type DatabaseClient struct {
	DB *gorm.DB
}

func DatabaseService() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := databases.Connect()
			d := DatabaseClient{DB: session}

			sqlDb, err := d.DB.DB()
			defer sqlDb.Close()
			if err != nil {
				logrus.Fatal(err)
			}

			//https://qiita.com/earl2/items/e2ae573128d077cf088e
			d.DB.Logger = d.DB.Logger.LogMode(logger.Info)

			c.Set("dbs", &d)
			if err := next(c); err != nil {
				return err
			}
			return nil
		}
	}
}
