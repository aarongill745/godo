package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aarongill745/godo/commands"
)

var GlobalTasksFile = "tasks.txt"

func main() {
	fmt.Println(os.Args[0])
	command := os.Args[1]
	task := strings.Join(os.Args[2:], "")
	fmt.Println(command, task)
	switch command {
	case "add":
		commands.Add(task)
	}

}

func init() {
	// Create new tasks file
	_, err := os.Stat(GlobalTasksFile)
	if os.IsNotExist(err) {
		os.Create(GlobalTasksFile)
	}
}
