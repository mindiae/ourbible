package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"

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

	htmlFileName := filepath.Join(APP_ROOT, "static/webview.tmpl")
	html, err := os.ReadFile(htmlFileName)
	if err != nil {
		fmt.Println("Error reading webview.tmpl file:", err)
		return
	}

	tmpl, err := template.New("main").Parse(string(html))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	modules, err := ModulesHandler()
	if err != nil {
		fmt.Println("Error getting modules", err)
	}
	modulesJson, _ := json.Marshal(modules)

	strings, err := GetStrings(db)
	if err != nil {
		fmt.Println("Error getting strings", err)
	}

	ints, err := GetInts(db)
	if err != nil {
		fmt.Println("Error getting numbers", err)
	}

	books, err := GetBooks(db)
	if err != nil {
		fmt.Println("Error getting books", err)
	}
	booksJson, _ := json.Marshal(books)

	booksTable, err := BooksHandler(strings["module"])
	if err != nil {
		fmt.Println("Error getting booksTable", err)
	}
	booksTableJson, _ := json.Marshal(booksTable)

	booksTable2 := []Book{}

	if strings["module2"] != "" {
		booksTable2, err = BooksHandler(strings["module2"])
		if err != nil {
			fmt.Println("Error getting booksTable", err)
		}
	}

	booksTable2Json, _ := json.Marshal(booksTable2)

	style := "<style>" + string(bibleviewer) + string(fontawesome) + "</style>"
	script := "<script>" + string(alpineSwipe) + string(alpine) + "</script>"

	data := struct {
		Style            string
		Javascript       string
		Module           string
		Module2          string
		BookNumber       int16
		Chapter          int16
		Verse            int16
		IsSystemDarkMode int16
		DarkMode         int16
		Books            string
		Modules          string
		BooksTable       string
		BooksTable2      string
	}{
		Style:            style,
		Javascript:       script,
		Module:           strings["module"],
		Module2:          strings["module2"],
		BookNumber:       ints["bookNumber"],
		Chapter:          ints["chapter"],
		Verse:            ints["verse"],
		IsSystemDarkMode: ints["isSystemDarkMode"],
		DarkMode:         ints["darkMode"],
		Books:            string(booksJson),
		Modules:          string(modulesJson),
		BooksTable:       string(booksTableJson),
		BooksTable2:      string(booksTable2Json),
	}

	var tpl bytes.Buffer

	if err := tmpl.Execute(&tpl, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	w := webview.New(true)
	w.SetTitle("OurBible")
	w.SetSize(800, 600, webview.HintNone)

	w.Bind("getBooks", BooksHandler)
	w.Bind("getModules", ModulesHandler)
	w.Bind("getChapters", ChapterHandler)
	w.Bind("setStringItem", func(key string, value string) error {
		return SetStringItem(db, key, value)
	})
	w.Bind("setNumberItem", func(key string, value int16) error {
		return SetIntItem(db, key, value)
	})
	w.Bind("setBookVerse", func(bookNumber int16, value int16) error {
		return SetBookVerse(db, bookNumber, value)
	})
	w.Bind("setBookChapter", func(bookNumber int16, value int16) error {
		return SetBookChapter(db, bookNumber, value)
	})

	w.SetHtml(tpl.String())

	w.Run()
}
