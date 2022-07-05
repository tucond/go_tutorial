//https://vivacocoa.jp/go/gin/gin_database.php

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorm/gorm"
	_ "github.com/mattn/go-sqlite3" //for gorm
)

func main() {
	router:=gin.Default()
	router.Static("styles","./styles")
	router.LoadHTMLGlob("./templates/*.html")
	dbInit()
	router.GET("/",getHandler)
	router.POST("/",postHandler)
	router.Run()
}

func getHandler(ctx *gin.Context) {
	todo := getAll()
	ctx.HTML(200,"index.html",gin.H{"todo":todo})
}

func postHandler(ctx *gin.Context)
	memo := ctx.PostForm("memo")
	create(memo)
	ctx.Redirect(302, "/")
}

func getAll() []Todo{ 
	db,err := gorm.Open("sqlite3","todo.sqlite3")
	if err != nil{
		panic("failed to connect db")
	}
	var todo []Todo
	dv.Find(&todo)
	return todo
}

func dbInit(){

}

func create(memo string){

}
