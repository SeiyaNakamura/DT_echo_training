package controller

import (
	"app/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

//articleの構造体
var articles []*models.Article

//articleのデータをtest.htmlに出力
func Test(c echo.Context) error {
	db := models.DatabaseConnection()
	articles := db.Find(&articles)
	return c.Render(http.StatusOK, "test", articles)
}
