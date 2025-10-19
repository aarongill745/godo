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
	var Reset = "\033[0m"
	var Red = "\033[31m"
	lineNumber := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("%s%d:%s %s\n", Red, lineNumber, Reset, scanner.Text())
		lineNumber++
	}
}
