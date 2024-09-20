package main

import (
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Verse struct {
	Text string
}

func ChapterHandler(module string, bookNumber int16, chapter int16) ([]string, error) {
	file := filepath.Join(APP_ROOT, "database", fmt.Sprintf("%s.SQLite3", module))
	if !fileExists(file) {
		file = filepath.Join(configPath, "database", fmt.Sprintf("%s.SQLite3", module))
		if !fileExists(file) {
			return []string{}, errors.New("file " + module + ".SQLite3" + " does not exist")
		}
	}

	var outputChapter []string

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return []string{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT text FROM verses WHERE book_number = ? AND chapter = ?", bookNumber, chapter)
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var verse Verse
		if err := rows.Scan(&verse.Text); err != nil {
			return []string{}, err
		}
		outputChapter = append(outputChapter, verse.Text)
	}
	if err := rows.Err(); err != nil {
		return []string{}, err
	}

	return outputChapter, nil
}
