package main

import (
  "fmt"
	"github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/labstack/gommon/log"
	"html/template"
	"io"
	"net/http"
  "github.com/kobakei/go-anond/models"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func GetArticles(c echo.Context) error {
  // TODO
  articles := []models.Article{}
  articles = append(articles, models.Article{"PPAP", "Pen Pineapple Apple Pen"})
  articles = append(articles, models.Article{"PPAP", "Pen Pineapple Apple Pen"})
  articles = append(articles, models.Article{"PPAP", "Pen Pineapple Apple Pen"})
	return c.Render(http.StatusOK, "articles/index", articles)
}

func NewArticle(c echo.Context) error {
	return c.Render(http.StatusOK, "articles/new", nil)
}

func GetArticle(c echo.Context) error {
  // TODO
	article := models.Article{"PPAP", "Pen Pineapple Apple Pen"}
	return c.Render(http.StatusOK, "articles/show", article)
}

func SaveArticle(c echo.Context) error {
  fmt.Println("Save")
	return c.Redirect(http.StatusMovedPermanently, "/articles/123")
}

func main() {
	e := echo.New()

  // ログの設定
  e.Debug = true
  e.Logger.SetLevel(log.DEBUG)

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

  // Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 起動
	e.Logger.Fatal(e.Start(":1323"))
}
