package main

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}
