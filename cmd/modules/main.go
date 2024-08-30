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

type InfoRow struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	directoryPath := "database"

	var modules []map[string]string

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

			rows, err := db.Query("SELECT * FROM info WHERE name != 'history_of_changes' ")
			if err != nil {
				log.Printf("Error querying database: %v", err)
				return err
			}
			defer rows.Close()

			info := make(map[string]string)

			for rows.Next() {
				var infoRow InfoRow
				if err := rows.Scan(&infoRow.Name, &infoRow.Value); err != nil {
					log.Printf("Error scanning rows: %v", err)
					return err
				}
				info[infoRow.Name] = infoRow.Value
			}
			info["name"] = moduleName

			if err := rows.Err(); err != nil {
				log.Printf("Error iterating rows: %v", err)
				return err
			}

			modules = append(modules, info)
		}
		return nil
	})
	if err != nil {
		log.Printf("Error reading directory: %v", err)
		return
	}

	jsonData, err := json.Marshal(modules)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return
	}

	err = os.WriteFile("static/bible-json/module", jsonData, 0o644)
	if err != nil {
		log.Printf("Error writing to file: %v", err)
		return
	}

	log.Println("JSON data written to output.json")
}
