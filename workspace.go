package main

import (
	"errors"
	"os"
	"path/filepath"
)

func GodoInitExists(currentDirectory string) bool {
	_, err := FindGodoInit(currentDirectory)
	return err != nil
}

func FindGodoInit(currentDirectory string) (string, error) {
	for {
		godoPath := filepath.Join(currentDirectory, ".godo")

		if info, err := os.Stat(godoPath); err == nil && info.IsDir() {
			return godoPath, nil
		}
		parentDir := filepath.Dir(currentDirectory)
		if parentDir == currentDirectory {
			return "", errors.New("error: workspace not found")
		}
		currentDirectory = parentDir
	}
}
