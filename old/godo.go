// The first thing that will get executed when running the program.
package main

import (
	"os"
	"strings"

	"github.com/aarongill745/godo/commands"
	"github.com/aarongill745/godo/utils"
)

var GlobalTasksFile = "tasks.txt"

func main() {
	command := os.Args[1]

	// Command handler
	switch command {
	case "add":
		body := strings.Join(os.Args[2:], " ")
		handleAdd(body)
	case "complete":
		body := os.Args[2]
		handleComplete(body)
	case "init":
		initApp()
	case "list":
		handleList()
	case "help":
		handleHelp(os.Args)
	case "edit":
		handleEdit(os.Args)
	}
}
func handleEdit(args []string) {
	commands.Edit(args[2], strings.Join(args[3:], " "), GlobalTasksFile)
}

func handleHelp(arguments []string) {
	utils.Help(arguments)
}

func handleList() {
	commands.List(GlobalTasksFile)
}

func handleAdd(task string) {
	commands.Add(task)
}

func handleComplete(line string) {
	commands.Complete(line, GlobalTasksFile)
}

func initApp() {
	utils.Init(GlobalTasksFile)
}
