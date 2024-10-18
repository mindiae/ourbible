package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

func BookHandler(module string, bookNumber int) (string, error) {
	file := filepath.Join(APP_ROOT, "database", fmt.Sprintf("%s.SQLite3", module))
	if !fileExists(file) {
		file = filepath.Join(configPath, fmt.Sprintf("%s.SQLite3", module))
		if !fileExists(file) {
			return "", errors.New("file " + module + ".SQLite3" + " does not exist")
		}
	}

	var book string

	db, err := sqlite3.Open(file)
	if err != nil {
		return "", err
	}
	defer db.Close()

	rows, err := db.Prepare("SELECT  chapter, verse, text FROM verses WHERE book_number = ? ", bookNumber)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	re := regexp.MustCompile(`(?P<word>[^><]+)<S>(?P<strong>\d+)<\/S>`)

	var lastVerse int
	for {
		hasRow, err := rows.Step()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if !hasRow {
			break
		}
		var chapter int
		var verse int
		var text string
		if err := rows.Scan(&chapter, &verse, &text); err != nil {
			return "", err
		}
		if lastVerse != 0 && verse == 1 {
			book = book + fmt.Sprintf(`<span x-init="maxVerses._%d_%d = %d"></span>`, bookNumber, chapter-1, lastVerse)
		}
		book = book + fmt.Sprintf(`<div id="verse_%d_%d_%d">`, bookNumber, chapter, verse)
		if chapter == 1 && verse == 1 {
			book = book + `<h1 x-text="bookstable.find((book) => book.book_number == book_number)?.long_name"></h1>`
		}
		if verse == 1 {
			book = book + fmt.Sprintf(`<h2 x-text="modules.find((mod)=>mod.name == module).chapter_string + ' ' + %d"></h2>`, chapter)
		}

		book = book + fmt.Sprintf(`<div>
      <span 
        class="float-left text-nowrap text-green-700 dark:text-green-500"
      >%d &nbsp;</span>
      <span>%s</span>
      </div>
      </div>`, verse, re.ReplaceAllString(text, `<span @click='S("${strong}")'>${word}</span>`))
		lastVerse = verse
	}

	return book, nil
}
