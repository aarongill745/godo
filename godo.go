package main

import (
	"fmt"
	"godo/godo/commands"
	"os"
	"strings"
)

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
