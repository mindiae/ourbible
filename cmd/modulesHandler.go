package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
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
			db, err := sqlite3.Open(path)
			if err != nil {
				return err
			}
			defer db.Close()

			rows, err := db.Prepare("SELECT * FROM info WHERE name != 'history_of_changes' ")
			if err != nil {
				return err
			}
			defer rows.Close()

			info := make(map[string]string)

			for {
				hasRow, err := rows.Step()
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
				if !hasRow {
					break
				}
				var infoRow InfoRow
				if err := rows.Scan(&infoRow.Name, &infoRow.Value); err != nil {
					return err
				}
				info[infoRow.Name] = infoRow.Value
			}
			info["name"] = moduleName

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
			db, err := sqlite3.Open(path)
			if err != nil {
				return err
			}
			defer db.Close()

			rows, err := db.Prepare("SELECT * FROM info WHERE name != 'history_of_changes' ")
			if err != nil {
				return err
			}
			defer rows.Close()

			info := make(map[string]string)

			for {
				hasRow, err := rows.Step()
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
				if !hasRow {
					break
				}
				var infoRow InfoRow
				if err := rows.Scan(&infoRow.Name, &infoRow.Value); err != nil {
					return err
				}
				info[infoRow.Name] = infoRow.Value
			}
			info["name"] = moduleName

			modules = append(modules, info)
		}
		return nil
	})
	if err != nil {
		return emptyReturnValue, err
	}

	return modules, nil
}
