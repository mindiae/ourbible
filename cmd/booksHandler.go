package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	BookNumber int16  `json:"book_number"`
	BookColor  string `json:"book_color"`
	ShortName  string `json:"short_name"`
	LongName   string `json:"long_name"`
}

func BooksHandler(c echo.Context) error {
	module := c.Param("module")

	file := filepath.Join(APP_ROOT, "database", fmt.Sprintf("%s.SQLite3", module))
	if !fileExists(file) {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database file does not exist\n" + file})
	}
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer db.Close()

	rows, err := db.Query("SELECT book_color, book_number, short_name, long_name FROM books")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.BookColor, &book.BookNumber, &book.ShortName, &book.LongName); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, books)
}
