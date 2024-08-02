package commands

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Edit(line string, newTask string, fileName string) {
	lineNumber, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal(err)
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	if lineNumber > len(lines) {
		fmt.Println("Invalid task")
		return
	}
	prevTask := lines[lineNumber-1]
	lines[lineNumber-1] = newTask
	err = os.WriteFile(fileName, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		confirmEdit(prevTask, newTask, line)
	}
}

func confirmEdit(prevTask string, newTask string, taskId string) {
	ColourReset := "\033[0m"
	ColourGreen := "\033[32m"
	ColourRed := "\033[31m"

	fmt.Println("------------")
	fmt.Println("")
	fmt.Printf("Task %s Edited Successfully", taskId)
	fmt.Printf("%s%s%s -> %s%s%s", ColourRed, prevTask, ColourReset, ColourGreen, newTask, ColourReset)
	fmt.Println("")
	fmt.Println("------------")
}
