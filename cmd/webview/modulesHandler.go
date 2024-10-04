package main

import (
	"database/sql"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type InfoRow struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func ModulesHandler() ([]map[string]string, error) {
	internalDbDir := filepath.Join(APP_ROOT, "database")
	configDbDir := configPath

	var emptyReturnValue []map[string]string
	var modules []map[string]string

	err := filepath.Walk(internalDbDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			moduleName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			db, err := sql.Open("sqlite3", path)
			if err != nil {
				return err
			}
			defer db.Close()

			rows, err := db.Query("SELECT * FROM info WHERE name != 'history_of_changes' ")
			if err != nil {
				return err
			}
			defer rows.Close()

			info := make(map[string]string)

			for rows.Next() {
				var infoRow InfoRow
				if err := rows.Scan(&infoRow.Name, &infoRow.Value); err != nil {
					return err
				}
				info[infoRow.Name] = infoRow.Value
			}
			info["name"] = moduleName

			if err := rows.Err(); err != nil {
				return err
			}

			modules = append(modules, info)
		}
		return nil
	})
	if err != nil {
		return emptyReturnValue, err
	}

	err = filepath.Walk(configDbDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(info.Name()) == ".SQLite3" {
			moduleName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			db, err := sql.Open("sqlite3", path)
			if err != nil {
				return err
			}
			defer db.Close()

			rows, err := db.Query("SELECT * FROM info WHERE name != 'history_of_changes' ")
			if err != nil {
				return err
			}
			defer rows.Close()

			info := make(map[string]string)

			for rows.Next() {
				var infoRow InfoRow
				if err := rows.Scan(&infoRow.Name, &infoRow.Value); err != nil {
					return err
				}
				info[infoRow.Name] = infoRow.Value
			}
			info["name"] = moduleName

			if err := rows.Err(); err != nil {
				return err
			}

			modules = append(modules, info)
		}
		return nil
	})
	if err != nil {
		return emptyReturnValue, err
	}

	return modules, nil
}
