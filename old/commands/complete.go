package commands

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Complete(lineToDelete string, fileName string) {
	lineNumber, err := strconv.Atoi(lineToDelete)
	if err != nil {
		log.Fatal(err)
		return
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	lines := strings.Split(string(content), "\n")
	if lineNumber > len(lines) {
		fmt.Println("Invalid task")
		return
	}

	lines = append(lines[:lineNumber-1], lines[lineNumber:]...)
	newContent := strings.Join(lines, "\n")

	err = os.WriteFile(fileName, []byte(newContent), 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
}
