package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	BookNumber int16  `json:"book_number"`
	BookColor  string `json:"book_color"`
	ShortName  string `json:"short_name"`
	LongName   string `json:"long_name"`
}

func main() {
	directoryPath := "database"

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error walking directory: %v", err)
			return err
		}
		if !info.IsDir() {
			moduleName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			db, err := sql.Open("sqlite3", path)
			if err != nil {
				log.Printf("Error opening database: %v", err)
				return err
			}
			defer db.Close()

			rows, err := db.Query("SELECT book_color, book_number, short_name, long_name FROM books")
			if err != nil {
				log.Printf("Error querying database: %v", err)
				return err
			}
			defer rows.Close()

			var books []Book

			for rows.Next() {
				var book Book
				if err := rows.Scan(&book.BookColor, &book.BookNumber, &book.ShortName, &book.LongName); err != nil {
					log.Printf("Error scanning rows: %v", err)
					return err
				}
				books = append(books, book)
			}

			if err := rows.Err(); err != nil {
				log.Printf("Error iterating rows: %v", err)
				return err
			}

			jsonData, err := json.Marshal(books)
			if err != nil {
				log.Printf("Error marshaling JSON: %v", err)
				return err
			}

			err = os.WriteFile("static/bible-json/book/"+moduleName, jsonData, 0o644)
			if err != nil {
				log.Printf("Error writing to file: %v", err)
				return err
			}

			log.Println("JSON data written to output.json")
		}
		return nil
	})
	if err != nil {
		log.Printf("Error reading directory: %v", err)
		return
	}
}
