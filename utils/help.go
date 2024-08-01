package utils

import (
	"fmt"
	"os"
)

// ANSI Colour codes
const (
	ColourBlue  = "\033[34m"
	ColourReset = "\033[0m"
	ColourGreen = "\033[32m"
)

// Help prints the usage information for the application.
func Help(arguments []string) {
	if len(arguments) == 2 {
		help()
	}

	if len(arguments) == 3 {
		switch arguments[2] {
		case "add":
			helpAdd()
		case "list":
			helpList()
		case "complete":
			helpComplete()
		case "help":
			helpHelp()
		case "init":
			helpInit()
		default:
			fmt.Printf("Unknown command: %s\n", os.Args[2])
			fmt.Println("Use 'godo help' for a list of available commands.")
		}
	}
}

func help() {
	fmt.Println()
	fmt.Println("Godo: A simple command-line todo application.")
	fmt.Println()
	fmt.Printf("%sUsage:%s\n", ColourBlue, ColourReset)
	fmt.Println()
	fmt.Println("  godo <command> [arguments]")
	fmt.Println()
	fmt.Printf("%sAvailable Commands:%s\n", ColourBlue, ColourReset)
	fmt.Println("  init			     Creates a new tasks file for you to start!")
	fmt.Println("  add <task>        Add a new task to your todo list")
	fmt.Println("  list              List all tasks")
	fmt.Println("  complete <id>     Mark a task as completed")
	fmt.Println("  help              Show help information")
	fmt.Println()
	fmt.Println("Use 'godo help <command>' for more information about a command.")
}

func helpAdd() {
	fmt.Println()
	fmt.Printf("%sAdd:%s A command to add a task that to your list", ColourRed, ColourReset)
	fmt.Println()
	fmt.Println()
	fmt.Printf("%sUsage:%s\n", ColourBlue, ColourReset)
	fmt.Println("  godo add [task]")
	fmt.Println()
}

func helpList() {
	fmt.Println()
	fmt.Printf("%sList:%s A command to list out all current outstanding tasks. This is also how you find out your task id.", ColourRed, ColourReset)
	fmt.Println()
	fmt.Println()
	fmt.Printf("%sUsage:%s\n", ColourBlue, ColourReset)
	fmt.Println("  godo list")
	fmt.Println()
}

func helpComplete() {
	fmt.Println()
	fmt.Printf("%sComplete:%s A command to complete a task that is currently in your list when given a task id", ColourRed, ColourReset)
	fmt.Println()
	fmt.Println()
	fmt.Printf("%sUsage:%s\n", ColourBlue, ColourReset)
	fmt.Println("  godo complete [taskId]")
	fmt.Println()
}

func helpHelp() {
	fmt.Println()
	fmt.Printf("%sHelp:%s A command to provide help on how to use this program.", ColourRed, ColourReset)
	fmt.Println()
	fmt.Println()
	fmt.Printf("%sUsage:%s\n", ColourBlue, ColourReset)
	fmt.Println("  godo help")
	fmt.Println("  godo help [anyCommand]")
	fmt.Println()
}

func helpInit() {
	fmt.Println()
	fmt.Printf("%sHelp:%s A command to initialise the godo app, creating a tasks.exe folder to store your tasks.", ColourRed, ColourReset)
	fmt.Println()
	fmt.Println()
	fmt.Printf("%sUsage:%s\n", ColourBlue, ColourReset)
	fmt.Println("  godo init")
	fmt.Println()
}
