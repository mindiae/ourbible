package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

// Book represents a book record
type Book struct {
	BookNumber int16  `json:"book_number"`
	BookColor  string `json:"book_color"`
	ShortName  string `json:"short_name"`
	LongName   string `json:"long_name"`
}

func booksHandler(c echo.Context) error {
	m := c.QueryParam("m")

	file := filepath.Join(APP_ROOT, "database", fmt.Sprintf("%s.SQLite3", m))
	if !fileExists(file) {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database file does not exist\n" + file})
	}
	// Open the SQLite database
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer db.Close()

	// Query the books table
	rows, err := db.Query("SELECT * FROM books") // Adjust the fields as per your schema
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()

	// Slice to hold the books
	var books []Book

	// Iterate through the rows
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.BookColor, &book.BookNumber, &book.ShortName, &book.LongName); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		books = append(books, book)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Return the list of books as JSON
	return c.JSON(http.StatusOK, books)
}
