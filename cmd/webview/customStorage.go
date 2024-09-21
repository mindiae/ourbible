package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

func getConfigPath(appName string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Print(err.Error())
	}

	if runtime.GOOS == "windows" {
		return filepath.Join(homeDir, "AppData", "Roaming", appName+".exe")
	}

	return filepath.Join(homeDir, ".local", "share", appName)
}

var (
	configPath  = getConfigPath("ourbible")
	storageFile = filepath.Join(configPath, "storage.sqlite3")
)

func GetStringItem(db *sql.DB, key string) (string, error) {
	var stringValue string
	err := db.QueryRow("SELECT value FROM strings WHERE key = ?", key).
		Scan(&stringValue)
	if err != nil {
		return "", err
	}

	return stringValue, nil
}

func SetStringItem(db *sql.DB, key string, value string) error {
	_, err := db.Exec("UPDATE strings SET value = ? WHERE key = ?", value, key)
	if err != nil {
		return err
	}

	return nil
}

func GetIntItem(db *sql.DB, key string) (int16, error) {
	var intValue int16
	err := db.QueryRow("SELECT value FROM numbers WHERE key = ?", key).
		Scan(&intValue)
	if err != nil {
		return 0, err
	}

	return intValue, nil
}

func SetIntItem(db *sql.DB, key string, value int16) error {
	_, err := db.Exec("UPDATE numbers SET value = ? WHERE key = ?", value, key)
	if err != nil {
		return err
	}

	return nil
}

func GetBooks(db *sql.DB) ([]map[string]int16, error) {
	books := []map[string]int16{}

	rows, err := db.Query("SELECT book_number, max_chapter, chapter, verse FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := map[string]int16{}
		var bookNumber int16
		var maxChapter int16
		var chapter int16
		var verse int16
		if err := rows.Scan(&bookNumber, &maxChapter, &chapter, &verse); err != nil {
			return nil, err
		}
		book["book_number"] = bookNumber
		book["max_chapter"] = maxChapter
		book["chapter"] = chapter
		book["verse"] = verse
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func SetBookVerse(db *sql.DB, bookNumber int16, value int16) error {
	_, err := db.Exec("UPDATE books SET verse = ? WHERE book_number = ?", value, bookNumber)
	if err != nil {
		return err
	}

	return nil
}

func SetBookChapter(db *sql.DB, bookNumber int16, value int16) error {
	_, err := db.Exec("UPDATE books SET chapter = ? WHERE book_number = ?", value, bookNumber)
	if err != nil {
		return err
	}

	return nil
}

func GetStrings(db *sql.DB) (map[string]string, error) {
	strings := map[string]string{}

	rows, err := db.Query("SELECT key, value FROM strings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var key string
		var value string
		if err := rows.Scan(&key, &value); err != nil {
			return nil, err
		}
		strings[key] = value
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return strings, nil
}

func GetInts(db *sql.DB) (map[string]int16, error) {
	ints := map[string]int16{}

	rows, err := db.Query("SELECT key, value FROM numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var key string
		var value int16
		if err := rows.Scan(&key, &value); err != nil {
			return nil, err
		}
		ints[key] = value
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ints, nil
}
