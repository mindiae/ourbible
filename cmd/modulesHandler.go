package main

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type InfoRow struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func ModulesHandler(c echo.Context) error {
	directoryPath := filepath.Join(APP_ROOT + "database")

	var modules []map[string]string

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			moduleName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			db, err := sql.Open("sqlite3", path)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			defer db.Close()

			rows, err := db.Query("SELECT * FROM info WHERE name != 'history_of_changes' ")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			defer rows.Close()

			info := make(map[string]string)

			for rows.Next() {
				var infoRow InfoRow
				if err := rows.Scan(&infoRow.Name, &infoRow.Value); err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
				}
				info[infoRow.Name] = infoRow.Value
			}
			info["name"] = moduleName

			if err := rows.Err(); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}

			modules = append(modules, info)
		}
		return nil
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error reading directory"})
	}

	return c.JSON(http.StatusOK, modules)
}
