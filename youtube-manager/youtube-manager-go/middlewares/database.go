package middlewares

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"youtube-manager-go/databases"
)

type DatabaseClient struct {
	DB *gorm.DB
}

func DatabaseService() echo.MiddlewareFunc{
	return func(next echo.HandlerFunc)echo.HandlerFunc{
		retrun func(c echo.Context)error{
			session,_:=databases.Connect()
			d:=DatabaseClient(DB:session)
			defer d.DB.Close()
			d.DB.LogMode(true)
			c.Set("dbs",&d)
			if err:=next(c);err!=nil{
				return err
			}
			return nil
		}
	}
}
