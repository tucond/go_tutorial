package main

import (
	"github.com/labstack/echo"
	"youtube-manager-go/routes"
)

func main() {
	e := echo.New()

	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
