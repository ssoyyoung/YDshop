package main

import (
	"github.com/labstack/echo"
	"github.com/ssoyyoung.p/YDshop/handler"
)

func main() {
	e := echo.New()

	e.GET("/getClothList", handler.GetClothList)

	e.Logger.Fatal(e.Start(":1324"))

}
