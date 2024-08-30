package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	webview "github.com/webview/webview_go"
)

var APP_ROOT = ""

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading it, using default value.")
	} else {
		if ROOT, exists := os.LookupEnv("ROOT"); exists {
			APP_ROOT = ROOT
		}
	}

	e := echo.New()

	if err != nil {
		log.Fatalf("Failed to initialize templates: %v", err)
	}

	e.GET("/bible-json/chapter/:module/:book/:chapter", ChapterHandler)
	e.GET("/bible-json/module", ModulesHandler)
	e.GET("/bible-json/book/:module", BooksHandler)
	e.Static("/", filepath.Join(APP_ROOT, "static"))

	// Start the Echo server in a goroutine
	go func() {
		e.Logger.Fatal(e.Start(":42069"))
	}()

	// Initialize the webview
	w := webview.New(false)
	w.SetTitle("OurBible")
	w.SetSize(800, 600, webview.HintNone)

	// Navigate to the correct port
	w.Navigate("http://localhost:42069") // Match the Echo server port

	w.Run()
}
