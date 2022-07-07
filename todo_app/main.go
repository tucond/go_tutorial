//https://vivacocoa.jp/go/gin/gin_database.php

package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
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

	router.GET("/", listHandler)
	router.POST("/new", createHandler)
	router.GET("delete/:id", deleteHandler)

	router.Run()
}

func listHandler(ctx *gin.Context) {
	todo := dbGetAll()
	ctx.HTML(200, "index.html", gin.H{"todo": todo})
}

func createHandler(ctx *gin.Context) {
	memo := ctx.PostForm("memo")
	dbCreate(memo)
	ctx.Redirect(302, "/")
}

func deleteHandler(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	dbDelete(id)
	ctx.Redirect(302, "/")
}

func dbGetAll() []Todo {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db¥n")
	}

	sqlDb, err := db.DB() //https://github.com/go-gorm/gorm/issues/3216
	defer sqlDb.Close()

	var todo []Todo
	db.Find(&todo)
	return todo
}

func dbInit() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db¥n")
	}

	sqlDb, err := db.DB()
	defer sqlDb.Close()

	db.AutoMigrate(&Todo{})
}

func dbCreate(memo string) {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db¥n")
	}

	sqlDb, err := db.DB()
	defer sqlDb.Close()

	db.Create(&Todo{Memo: memo})
}

func dbDelete(id int) {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db¥n")
	}

	sqlDb, err := db.DB()
	defer sqlDb.Close()

	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
}
