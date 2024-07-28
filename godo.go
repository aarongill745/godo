// The first thing that will get executed when running the program.
package main

import (
	"os"
	"strings"

	"github.com/aarongill745/godo/commands"
)

var GlobalTasksFile = "tasks.txt"

func main() {
	command := os.Args[1]
	body := strings.Join(os.Args[2:], "")

	// Command handler
	switch command {
	case "add":
		handleAdd(body)
	case "complete":
		handleComplete(body)
	}
}

func handleAdd(task string) {
	commands.Add(task)
}

func handleComplete(line string) {
	commands.Complete(line, GlobalTasksFile)
}

func init() {
	// Create new tasks file
	_, err := os.Stat(GlobalTasksFile)
	if os.IsNotExist(err) {
		os.Create(GlobalTasksFile)
	}
}
