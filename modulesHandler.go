package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

// ModulesHandler handles the request to list files without extensions
func modulesHandler(c echo.Context) error {
	// Specify the directory path
	directoryPath := APP_ROOT + "/database"

	// Slice to hold filenames without extensions
	var filenamesWithoutExtensions []string

	// Walk through the directory
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if it's a file
		if !info.IsDir() {
			// Get the filename without extension
			filename := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			filenamesWithoutExtensions = append(filenamesWithoutExtensions, filename)
		}
		return nil
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error reading directory"})
	}

	// Convert the slice to JSON
	return c.JSON(http.StatusOK, filenamesWithoutExtensions)
}
