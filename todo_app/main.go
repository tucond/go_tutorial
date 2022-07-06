//https://vivacocoa.jp/go/gin/gin_database.php

package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Memo string
}

func main() {
	router := gin.Default()
	router.Static("styles", "./styles")
	router.LoadHTMLGlob("./templates/*.html")
	dbInit()
	router.GET("/", getHandler)
	router.POST("/new", postHandler)
	router.Run()
}

func getHandler(ctx *gin.Context) {
	todo := getAll()
	ctx.HTML(200, "index.html", gin.H{"todo": todo})
}

func postHandler(ctx *gin.Context) {
	memo := ctx.PostForm("memo")
	create(memo)
	ctx.Redirect(302, "/")
}

func getAll() []Todo {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db¥n")
	}
	var todo []Todo
	db.Find(&todo)
	return todo
}

func dbInit() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db¥n")
	}
	db.AutoMigrate(&Todo{})
}

func create(memo string) {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db¥n")
	}
	db.Create(&Todo{Memo: memo})
}
