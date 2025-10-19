package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Task struct {
	Task string
}

// Need uppercase to export
func Add(task string) {
	fmt.Println(task)
	response, err := checkIsTaskDuplicate(task, "tasks.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	if response {
		fmt.Println("The task already exists")
		return
	}

	file, err := os.OpenFile("tasks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(task + "\n")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Added new task")
}

func checkIsTaskDuplicate(task string, fileName string) (bool, error) {
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
