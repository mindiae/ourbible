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
	Text string
}

func ChapterHandler(c echo.Context) error {
	module := c.Param("module")

	file := filepath.Join(APP_ROOT, "database", fmt.Sprintf("%s.SQLite3", module))
	if !fileExists(file) {
		return c.String(http.StatusInternalServerError, "Database file does not exist\n"+file)
	}

	bookNumber := c.Param("book")
	chapter := c.Param("chapter")

	var outputChapter []string

	db, err := sql.Open("sqlite3", filepath.Join(APP_ROOT, "database", fmt.Sprintf("%s.SQLite3", module)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer db.Close()

	rows, err := db.Query("SELECT text FROM verses WHERE book_number = ? AND chapter = ?", bookNumber, chapter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()

	for rows.Next() {
		var verse Verse
		if err := rows.Scan(&verse.Text); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		outputChapter = append(outputChapter, verse.Text)
	}
	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, outputChapter)
}
