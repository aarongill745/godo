package utils

import (
	"fmt"
)

// ANSI color codes
const (
	ColorBlue  = "\033[34m"
	ColorReset = "\033[0m"
)

// Help prints the usage information for the application.
func Help() {
	fmt.Println()
	fmt.Println("Godo: A simple command-line todo application.")
	fmt.Println()
	fmt.Printf("%sUsage:%s\n", ColorBlue, ColorReset)
	fmt.Println("  godo <command> [arguments]")
	fmt.Println()
	fmt.Printf("%sAvailable Commands:%s\n", ColorBlue, ColorReset)
	fmt.Println("  init			     Creates a new tasks file for you to start!")
	fmt.Println("  add <task>        Add a new task to your todo list")
	fmt.Println("  list              List all tasks")
	fmt.Println("  complete <id>     Mark a task as completed")
	fmt.Println("  help              Show help information")
	fmt.Println()
	fmt.Println("Use 'godo help <command>' for more information about a command.")
}
