package storage

import (
	"github.com/aarongill745/godo/models"
)

// ReadTasks reads all tasks from tasks.jsonl
// Returns empty slice if file doesn't exist or is empty
// Each line in the file should be a JSON-encoded Task
func ReadTasks() ([]models.Task, error) {
	// TODO: Get workspace path using GetWorkspacePath()
	// TODO: Open tasks.jsonl file
	// TODO: Use bufio.Scanner to read line by line
	// TODO: For each line:
	//       - Skip empty lines
	//       - json.Unmarshal into Task struct
	//       - If unmarshal fails, log warning and continue (corrupted line)
	//       - Add task to slice
	// TODO: Return slice of tasks
	panic("not implemented")
}

// WriteTasks overwrites tasks.jsonl with the provided tasks
// This is used for operations like deleting or editing tasks
// Each task is written as a JSON line
func WriteTasks(tasks []models.Task) error {
	// TODO: Get workspace path
	// TODO: Create/truncate tasks.jsonl file
	// TODO: For each task:
	//       - Marshal to JSON using json.Marshal()
	//       - Write JSON line + "\n" to file
	// TODO: Return error if any write fails
	panic("not implemented")
}

// AppendTask adds a single task to the end of tasks.jsonl
// This is faster than WriteTasks for adding new tasks
// The task must already have its ID set
func AppendTask(task models.Task) error {
	// TODO: Get workspace path
	// TODO: Open tasks.jsonl in append mode (os.O_APPEND | os.O_CREATE | os.O_WRONLY)
	// TODO: Marshal task to JSON
	// TODO: Write JSON + "\n" to file
	// TODO: Close file and return error if any
	panic("not implemented")
}

// GetTaskByID finds and returns a task with the given ID
// Returns nil and error if task not found
func GetTaskByID(id int) (*models.Task, error) {
	// TODO: Call ReadTasks() to get all tasks
	// TODO: Loop through tasks to find one with matching ID
	// TODO: If found, return pointer to task
	// TODO: If not found, return nil and error: "task not found"
	panic("not implemented")
}

// DeleteTask removes a task with the given ID from tasks.jsonl
// Returns error if task not found
func DeleteTask(id int) error {
	// TODO: Call ReadTasks() to get all tasks
	// TODO: Create new slice without the task with matching ID
	// TODO: If no task was removed, return error: "task not found"
	// TODO: Call WriteTasks() with the new slice
	panic("not implemented")
}

// NextTaskID returns the next available task ID
// Scans all existing tasks and returns max ID + 1
// Returns 1 if no tasks exist
func NextTaskID() (int, error) {
	// TODO: Call ReadTasks() to get all tasks
	// TODO: Loop through tasks to find max ID
	// TODO: Return maxID + 1 (or 1 if no tasks exist)
	panic("not implemented")
}

// UpdateTask updates an existing task with new data
// Finds task by ID and replaces it
func UpdateTask(task models.Task) error {
	// TODO: Call ReadTasks() to get all tasks
	// TODO: Find task with matching ID and replace it
	// TODO: If not found, return error: "task not found"
	// TODO: Call WriteTasks() with updated slice
	panic("not implemented")
}
