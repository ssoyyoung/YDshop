package handler

import (
	"net/http"

	"github.com/labstack/echo"
	mongoDB "github.com/ssoyyoung.p/clothForYD/mongoDB"
)

// GetClothList func
func GetClothList(c echo.Context) error {

	clothes := mongoDB.GetClothList()

	return c.JSON(http.StatusOK, clothes)

}
