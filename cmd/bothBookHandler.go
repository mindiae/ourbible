package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

func BothBookHandler(module string, module2 string, bookNumber int) (string, error) {
	file := filepath.Join(APP_ROOT, "database", fmt.Sprintf("%s.SQLite3", module))
	if !fileExists(file) {
		file = filepath.Join(configPath, fmt.Sprintf("%s.SQLite3", module))
		if !fileExists(file) {
			return "", errors.New("file " + module + ".SQLite3" + " does not exist")
		}
	}
	file2 := filepath.Join(APP_ROOT, "database", fmt.Sprintf("%s.SQLite3", module2))
	if !fileExists(file) {
		file2 = filepath.Join(configPath, fmt.Sprintf("%s.SQLite3", module2))
		if !fileExists(file) {
			return "", errors.New("file " + module2 + ".SQLite3" + " does not exist")
		}
	}

	var book string

	db, err := sqlite3.Open(file)
	if err != nil {
		return "", err
	}
	defer db.Close()

	attachQuery := fmt.Sprintf("ATTACH DATABASE '%s' AS db2", file2)
	err = db.Exec(attachQuery)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	rows, err := db.Prepare(`
    SELECT a.chapter, a.verse, a.text, b.text
        FROM verses AS a 
        INNER JOIN db2.verses AS b
          ON  a.book_number = b.book_number
          AND a.chapter     = b.chapter
          AND a.verse       = b.verse
        WHERE a.book_number = ?`, bookNumber)
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
		var text2 string
		if err := rows.Scan(&chapter, &verse, &text, &text2); err != nil {
			return "", err
		}
		if lastVerse != 0 && verse == 1 {
			book = book + fmt.Sprintf(`<span x-init="maxVerses._%d_%d = %d"></span>`, bookNumber, chapter-1, lastVerse)
		}
		book = book + fmt.Sprintf(`<div id="verse_%d_%d_%d">`, bookNumber, chapter, verse)
		if chapter == 1 && verse == 1 {
			book = book + `<div class="flex gap-2">
      <h1 class="basis-1/2" x-text="bookstable.find((book) => book.book_number == book_number)?.long_name"></h1>
      <h1 class="basis-1/2" x-text="bookstable2.find((book) => book.book_number == book_number)?.long_name"></h1>
      </div>`
		}
		if verse == 1 {
			book = book + fmt.Sprintf(`<div class="flex gap-2">
        <h2 class="basis-1/2" x-text="modules.find((mod)=>mod.name == module).chapter_string + ' ' + %d"></h2>
        <h2 class="basis-1/2" x-text="modules.find((mod)=>mod.name == module2)?.chapter_string + ' ' + %d"></h2>
        </div>`, chapter, chapter)
		}

		book = book + fmt.Sprintf(`<div class="flex gap-2">
      <span class="basis-1/2">
      <span 
        class="float-left text-nowrap text-green-700 dark:text-green-500"
      >%d &nbsp;</span>
      <span>%s</span>
      </span>
      <span class="basis-1/2">%s</span>
      </div>
      </div>
      `, verse, text, text2)
		lastVerse = verse
	}

	return re.ReplaceAllString(book, `<span @click='S("${strong}")'>${word}</span>`), nil
	// return book, nil
}
