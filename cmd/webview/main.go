package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
	webview "github.com/webview/webview_go"
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

var configPath = getConfigPath("ourbible")

func copyFile(src string, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

var APP_ROOT = ""

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func main() {
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Print("Error getting executable path:", err)
		return
	}

	if executablePath == "/usr/local/bin/ourbible" {
		APP_ROOT = "/usr/local/share/ourbible"
	}

	w := webview.New(true)
	w.SetTitle("OurBible")
	w.SetSize(800, 600, webview.HintNone)

	w.Bind("getBooks", BooksHandler)
	w.Bind("getModules", ModulesHandler)
	w.Bind("getChapters", ChapterHandler)

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
	}

	var appFullPath string
	if APP_ROOT == "" {
		appFullPath = cwd
	} else {
		appFullPath = APP_ROOT
	}
	w.Navigate("file://" + filepath.Join(appFullPath, "static/webview.html"))

	w.Run()
}
