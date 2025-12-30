package storage

import (
	"os"
)

// GetWorkspacePath returns the path to the global .godo directory
// Location: ~/.godo/ (cross-platform)
// Returns error if home directory cannot be determined
func GetWorkspacePath() (string, error) {
	// TODO: Use os.UserHomeDir() to get home directory
	// TODO: Join with ".godo" using filepath.Join()
	// TODO: Return the path
	panic("not implemented")
}

// WorkspaceExists checks if ~/.godo/ directory exists
func WorkspaceExists() (bool, error) {
	// TODO: Get workspace path using GetWorkspacePath()
	// TODO: Use os.Stat() to check if directory exists
	// TODO: Return true if exists and is a directory, false otherwise
	panic("not implemented")
}

// InitializeWorkspace creates the ~/.godo/ directory and empty files
// Creates: tasks.jsonl, completed.jsonl, tags.json, projects.json
// Returns error if workspace already exists or creation fails
func InitializeWorkspace() error {
	// TODO: Check if workspace already exists using WorkspaceExists()
	// TODO: If exists, return error: "workspace already initialized"
	// TODO: Get workspace path
	// TODO: Create .godo directory with os.MkdirAll()
	// TODO: Create empty files:
	//       - tasks.jsonl (empty file)
	//       - completed.jsonl (empty file)
	//       - tags.json (with {"tags": []})
	//       - projects.json (with {"projects": []})
	// TODO: Return nil on success
	panic("not implemented")
}

// Helper function to create an empty file at the given path
func createEmptyFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	return file.Close()
}

// Helper function to create a JSON file with initial content
func createJSONFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
