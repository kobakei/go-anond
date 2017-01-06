package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func GetArticles(c echo.Context) error {
  return c.String(http.StatusOK, "Index page")
}

func NewArticle(c echo.Context) error {
  return c.String(http.StatusOK, "New page")
}

func GetArticle(c echo.Context) error {
  return c.String(http.StatusOK, "Get page")
}

func SaveArticle(c echo.Context) error {
  return c.String(http.StatusOK, "Save page")
}

func main() {
	e := echo.New()
	e.GET("/", GetArticles)
  e.GET("/articles", GetArticles)
  e.GET("/articles/new", NewArticle)
  e.GET("/articles/:id", GetArticle)
  e.POST("/articles", SaveArticle)
	e.Logger.Fatal(e.Start(":1323"))
}
