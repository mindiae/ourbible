package main

import (
	"html/template"
	"io"
	"os"
	//"log"
	"net/http"

	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
)

const APP_ROOT = "/usr/local/share/ourbible"

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() (*Templates, error) {
	templates := template.Must(template.ParseGlob(APP_ROOT + "/views/*.html"))
	return &Templates{templates: templates}, nil
}

type Count struct {
	Count int
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func main() {
	e := echo.New()
	e.Static("/", APP_ROOT+"/assets")
	//	e.Use(middleware.Logger())

	//_, err := newTemplate()
	//	if err != nil {
	//	log.Fatal(err)
	//}

	// count := Count{Count: 0}
	e.Renderer, _ = newTemplate()

	e.GET("/", func(c echo.Context) error {
		// count.Count++
		return c.Render(http.StatusOK, "index", nil)
	})

	e.GET("/api-html/chapter", chapterHandler)

	e.GET("/api-json/modules", modulesHandler)

	e.GET("/api-json/books", booksHandler)

	e.Logger.Fatal(e.Start(":42069"))
}
