package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Task string
}

// Need uppercase to export
func Add(task string) {
	CheckIsTaskDuplicate(task, "tasks.txt")
	err := os.WriteFile("tasks.txt", []byte(task), 0777)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Godo: New task added")
	}
}

func CheckIsTaskDuplicate(task string, fileName string) (bool, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return false, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), task) {
			fmt.Printf("This task already exists, ID: %v", line)
			return true, nil
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}
	return false, nil
}
