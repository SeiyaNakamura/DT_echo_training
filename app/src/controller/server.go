package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init() {

	list, err := template.New("t").ParseGlob("template/*.html")
	t := &Template{
		templates: template.Must(list, err),
	}

	e := echo.New()
	e.Renderer = t
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//router
	e.GET("/", hello)
	//articleテーブルのデータを出力
	e.GET("/test", Test)

	// start server
	e.Logger.Fatal(e.Start(":8080"))

}
