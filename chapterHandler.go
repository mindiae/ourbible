package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
)

type Verse struct {
	Verse int8
	Text1 string
	Text2 string
}

type SingleVerse struct {
	Verse int8
	Text  string
}

type PageData struct {
	Verses               []Verse
	IsSecondModuleExists bool
}

func chapterHandler(c echo.Context) error {
	m := c.QueryParam("m")

	file := filepath.Join(APP_ROOT, "database", fmt.Sprintf("%s.SQLite3", m))
	if !fileExists(file) {
		return c.String(http.StatusInternalServerError, "Database file does not exist\n"+file)
	}

	m2 := c.QueryParam("m2")
	file = filepath.Join(APP_ROOT, "database", fmt.Sprintf("%s.SQLite3", m2))
	if m2 != "" && !fileExists(file) {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database file does not exist\n" + file})
	}
	bookNumber := c.QueryParam("b")
	chapter := c.QueryParam("c")

	modules := []string{m}
	if m2 != "" {
		modules = append(modules, m2)
	}

	var combinedResults []Verse

	// Create a map to store verses for the current module
	var firstResults []SingleVerse
	var secondResults []SingleVerse

	for index, module := range modules {
		db, err := sql.Open("sqlite3", filepath.Join(APP_ROOT, "database", fmt.Sprintf("%s.SQLite3", module)))
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		defer db.Close()

		rows, err := db.Query("SELECT verse, text FROM verses WHERE book_number = ? AND chapter = ?", bookNumber, chapter)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var verse int8
			var text string
			if err := rows.Scan(&verse, &text); err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			if index == 0 {
				firstResults = append(firstResults, SingleVerse{
					Verse: verse,
					Text:  text,
				})
			} else {
				secondResults = append(secondResults, SingleVerse{
					Verse: verse,
					Text:  text,
				})
			}
		}
		if err := rows.Err(); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

	}

	// Determine the maximum length to iterate
	maxLength := len(firstResults)
	if len(secondResults) > maxLength {
		maxLength = len(secondResults)
	}

	// Populate combinedResults
	for i := 0; i < maxLength; i++ {
		var verse Verse
		verse.Verse = int8(i + 1)
		if i < len(firstResults) {
			verse.Text1 = firstResults[i].Text
		}
		if i < len(secondResults) {
			verse.Text2 = secondResults[i].Text
		}
		combinedResults = append(combinedResults, verse)
	}

	data := PageData{
		Verses:               combinedResults,
		IsSecondModuleExists: m2 != "",
	}
	return c.Render(http.StatusOK, "verses", data)
}
