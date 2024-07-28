// The first thing that will get executed when running the program.
package main

import (
	"fmt"
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
	case "init":
		initApp()
	}

}

func handleAdd(task string) {
	commands.Add(task)
}

func handleComplete(line string) {
	commands.Complete(line, GlobalTasksFile)
}

// init is a reserved keyword
func initApp() {
	// Create new tasks file
	_, err := os.Stat(GlobalTasksFile)
	if os.IsNotExist(err) {
		os.Create(GlobalTasksFile)
		return
	}

	var response string
	fmt.Println("You have already done this before, are you sure about this? Y / N")
	for {
		fmt.Scanln(&response)
		response = strings.ToLower(response)
		// Input is valid
		if response == "y" || response == "n" {
			break
		} else {
			fmt.Println("Invalid input: Y / N")
		}
	}
}
