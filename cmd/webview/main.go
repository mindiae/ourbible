package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	webview "github.com/webview/webview_go"
)

func copyFile(src string, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

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

	if !fileExists(storageFile) {
		err := os.MkdirAll(filepath.Join(configPath, "database"), 0o755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}

		err = copyFile(filepath.Join(APP_ROOT, "storage.sqlite3"), storageFile)
		if err != nil {
			fmt.Printf("Error copying file: %v\n", err)
			return
		}
	}

	db, err := sql.Open("sqlite3", storageFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	head := `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />

    <style>
  `

	alpineSwipeFileName := filepath.Join(APP_ROOT, "static/js/alpinejs-swipe.js")
	alpineSwipe, err := os.ReadFile(alpineSwipeFileName)
	if err != nil {
		fmt.Println("Error reading alpine-swipe.js file:", err)
		return
	}

	alpineFileName := filepath.Join(APP_ROOT, "static/js/alpine.js")
	alpine, err := os.ReadFile(alpineFileName)
	if err != nil {
		fmt.Println("Error reading alpine.js file:", err)
		return
	}

	bibleviewerFileName := filepath.Join(APP_ROOT, "static/css/bibleviewer.css")
	bibleviewer, err := os.ReadFile(bibleviewerFileName)
	if err != nil {
		fmt.Println("Error reading bibleviewer.css file:", err)
		return
	}

	fontawesomeFileName := filepath.Join(APP_ROOT, "static/css/fontawesome.css")
	fontawesome, err := os.ReadFile(fontawesomeFileName)
	if err != nil {
		fmt.Println("Error reading fontawesome.css file:", err)
		return
	}

	htmlFileName := filepath.Join(APP_ROOT, "static/webview.html")
	html, err := os.ReadFile(htmlFileName)
	if err != nil {
		fmt.Println("Error reading webview.html file:", err)
		return
	}

	w := webview.New(false)
	w.SetTitle("OurBible")
	w.SetSize(800, 600, webview.HintNone)

	w.Bind("getBooks", BooksHandler)
	w.Bind("getModules", ModulesHandler)
	w.Bind("getChapters", ChapterHandler)
	w.Bind("getStringItem", func(key string) (string, error) {
		return GetStringItem(db, key)
	})
	w.Bind("setStringItem", func(key string, value string) error {
		return SetStringItem(db, key, value)
	})
	w.Bind("getNumberItem", func(key string) (int16, error) {
		return GetIntItem(db, key)
	})
	w.Bind("setNumberItem", func(key string, value int16) error {
		return SetIntItem(db, key, value)
	})
	w.Bind("getBooks", func() ([]map[string]int16, error) {
		return GetBooks(db)
	})
	w.Bind("setBookVerse", func(bookNumber int16, value int16) error {
		return SetBookVerse(db, bookNumber, value)
	})
	w.Bind("setBookChapter", func(bookNumber int16, value int16) error {
		return SetBookChapter(db, bookNumber, value)
	})

	w.SetHtml(head +
		string(bibleviewer) +
		string(fontawesome) +
		`</style></head>` +
		string(html) +
		`<script>` +
		string(alpineSwipe) +
		string(alpine) +
		`</script></html>`)

	w.Run()
}
