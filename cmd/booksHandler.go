package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

type Book struct {
	BookNumber int    `json:"book_number"`
	BookColor  string `json:"book_color"`
	ShortName  string `json:"short_name"`
	LongName   string `json:"long_name"`
}

func BooksHandler(module string) ([]Book, error) {
	file := filepath.Join(APP_ROOT, "database", fmt.Sprintf("%s.SQLite3", module))
	if !fileExists(file) {
		file = filepath.Join(configPath, fmt.Sprintf("%s.SQLite3", module))
		if !fileExists(file) {
			return []Book{}, errors.New("file " + module + ".SQLite3" + " does not exist")
		}
	}
	var emptyReturnValue []Book

	db, err := sqlite3.Open(file)
	if err != nil {
		return emptyReturnValue, err
	}
	defer db.Close()

	rows, err := db.Prepare("SELECT book_color, book_number, short_name, long_name FROM books")
	if err != nil {
		return emptyReturnValue, err
	}
	defer rows.Close()

	var books []Book

	for {
		hasRow, err := rows.Step()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if !hasRow {
			break
		}
		var book Book
		if err := rows.Scan(&book.BookColor, &book.BookNumber, &book.ShortName, &book.LongName); err != nil {
			return emptyReturnValue, err
		}
		books = append(books, book)
	}

	return books, nil
}
