package main

import (
	"fmt"
	"godo/godo/commands"
	"os"
	"strings"
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

// Creates a tasks file if it doesnt exist
func init() {
	_, err := os.Stat(GlobalTasksFile)
	if os.IsNotExist(err) {
		os.Create(GlobalTasksFile)
	}
}
