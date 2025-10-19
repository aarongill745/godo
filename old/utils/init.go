package utils

import (
	"fmt"
	"os"
	"strings"
)

func Init(GlobalTasksFile string) {
	_, err := os.Stat(GlobalTasksFile)

	// Create new tasks file
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
	return
}
