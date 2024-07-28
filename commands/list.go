package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Lists all current tasks
func List(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lineNumber := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("%d: %s\n", lineNumber, scanner.Text())
		lineNumber++
	}
}
