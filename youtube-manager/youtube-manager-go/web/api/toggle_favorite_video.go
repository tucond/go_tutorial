package api

import (
	"errors"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
	"youtube-manager-go/middlewares"
	"youtube-manager-go/models"
)

type ToggleFavoriteVideoResponse struct {
	VideoId   string `json:"video_id"`
	IsFavorte bool   `json:"is_favorite"`
}

func ToggleFavoriteVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		videoId := c.Param("id")
		token := c.Get("auth").(*auth.Token)
		user := models.User{}

		db := dbs.DB.Table("users").
			Where(models.User{UID: token.UID}).First(&user)
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {

			user = models.User{UID: token.UID}
			dbs.DB.Create(&user)
		}

		favorite := models.Favorite{}
		isFavorte := false

		db = dbs.DB.Table("favorites").
			Where(models.Favorite{UserId: user.ID, VideoId: videoId}).First(&favorite)
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			favorite = models.Favorite{UserId: user.ID, VideoId: videoId}
			dbs.DB.Create(&favorite)
			isFavorte = true
		} else {
			dbs.DB.Delete(&favorite)
		}

		res := ToggleFavoriteVideoResponse{
			VideoId:   videoId,
			IsFavorte: isFavorte,
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
