//https://vivacocoa.jp/go/gin/gin_firststep.php

package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router:=gin.Default()
	router.Static("styles","./styles")
	router.LoadHTMLGlob("./templates/*.html")
	router.GET("/",handler)
	router.Run()
}

func handler(ctx *gin.Context) {

	now:=time.Now()
	ctx.HTML(200,"index.html",gin.H{"hour":now.Hour(),"minute":now.Minute()})
}

