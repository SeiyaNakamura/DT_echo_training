package controller

import (
	"app/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

var articles []*models.Article

func Test(c echo.Context) error {
	db := models.DatabaseConnection()
	articles := db.Find(&articles)
	return c.Render(http.StatusOK, "test", articles)
}

////articleテーブルの取り出し
//db := models.DatabaseConnection()
//articles := db.Find(&articles)
