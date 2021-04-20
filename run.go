package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func getSupportedImages() []string {
	return []string{".png", ".jpg", ".jpeg", ".gif"}
}

func getSupportedDocuments() []string {
	return []string{".pdf", ".docx", ".doc", ".txt", ".log", ".odt", ".tex", ".xlsx", ".epub"}
}

func getSupportedArchives() []string {
	return []string{".7z", ".rar", ".tar.gz", ".z", ".zip", ".deb", ".tar.gz", ".gz", ".tgz"}
}

func getSupportedVideos() []string {
	return []string{".mp4", ".webm"}
}

func createDirectory(path string) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(path, 0755)

		if errDir != nil {
			log.Fatal(errDir)
		}
	}
}

func contains(items []string, elem string) (int, bool) {
	for i, item := range items {
		if item == elem {
			return i, true
		}
	}
	return -1, false
}

func moveFile(file string, folder string, types []string) {
	_, isSupported := contains(types, filepath.Ext(file))

	if isSupported {
		err := os.Rename(file, filepath.Join(file, folder, filepath.Base(file)))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[%s] Moved to [%s]\n", file, folder)
	}
}

func organize(files []string) {
	for _, file := range files {
		// Right now just try to move file to all folders
		// And it will be only moved to the right directory
		moveFile(file, "../Images", getSupportedImages())
		moveFile(file, "../Documents", getSupportedDocuments())
		moveFile(file, "../Archives", getSupportedArchives())
		moveFile(file, "../Videos", getSupportedVideos())
	}
}

func getFiles(dir string) []string {
	// Check if this specified directory exists or not
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatal("Directory does not exist")
	}

	// Read directory's files
	files, err := ioutil.ReadDir(dir)

	// In case there was an error while reading the directory
	if err != nil {
		log.Fatal(err)
	}

	var _files []string

	// Now loop over directory file is of type: os.FileInfo
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(dir, file.Name())
			_files = append(_files, filePath)
		}
	}
	return _files
}

func main() {
	// Get all arguments
	args := os.Args[1:]

	// Check if user specified any arguments
	if len(args) == 0 {
		log.Fatal("You must provide directory path")
	}

	// Get first argument which should be the directory path
	dir := args[0]

	// Initialize our own directories
	directories := []string{"images", "Documents", "Archives", "Videos"}

	for _, directory := range directories {
		createDirectory(filepath.Join(dir, directory))
	}

	// Get all files under this directory
	files := getFiles(dir)

	// Organize the files
	organize(files)
}
