package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	webview "github.com/webview/webview_go"
)

var APP_ROOT = ""

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func main() {
	executablePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Error getting executable path:", err)
		return
	}

	if executablePath == "/usr/local/bin/ourbible" {
		APP_ROOT = "/usr/local/share/ourbible"
	}

	e := echo.New()

	e.GET("/bible-json/chapter/:module/:book/:chapter", ChapterHandler)
	e.GET("/bible-json/module", ModulesHandler)
	e.GET("/bible-json/book/:module", BooksHandler)
	e.Static("/", filepath.Join(APP_ROOT, "static"))

	go func() {
		e.Logger.Fatal(e.Start(":42069"))
	}()

	w := webview.New(false)
	w.SetTitle("OurBible")
	w.SetSize(800, 600, webview.HintNone)

	w.Navigate("http://localhost:42069")
	w.Run()
}
