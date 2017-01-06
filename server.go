package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func GetArticles(c echo.Context) error {
	return c.Render(http.StatusOK, "articles/index", "World")
}

func NewArticle(c echo.Context) error {
	return c.Render(http.StatusOK, "articles/new", "World")
}

func GetArticle(c echo.Context) error {
	return c.Render(http.StatusOK, "articles/show", "World")
}

func SaveArticle(c echo.Context) error {
	return c.Render(http.StatusOK, "articles/create", "World")
}

func main() {
	e := echo.New()

	// テンプレートの設定
	t := &Template{
		templates: template.Must(template.ParseGlob("views/**/*.html")),
	}
	e.Renderer = t

	// ルーティング
	e.GET("/", GetArticles)
	e.GET("/articles", GetArticles)
	e.GET("/articles/new", NewArticle)
	e.GET("/articles/:id", GetArticle)
	e.POST("/articles", SaveArticle)

	// 起動
	e.Logger.Fatal(e.Start(":1323"))
}
