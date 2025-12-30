package storage

// ReadProjects reads all registered projects from projects.json
// File format: {"projects": ["backend", "frontend", "mobile"]}
// Returns empty slice if file doesn't exist or is empty
func ReadProjects() ([]string, error) {
	// TODO: Get workspace path
	// TODO: Read projects.json file
	// TODO: Unmarshal into ProjectRegistry struct
	// TODO: Return the Projects slice
	// TODO: If file doesn't exist, return empty slice (not an error)
	panic("not implemented")
}

// WriteProjects overwrites projects.json with the provided projects
func WriteProjects(projects []string) error {
	// TODO: Get workspace path
	// TODO: Create ProjectRegistry struct with projects
	// TODO: Marshal to JSON with indentation (json.MarshalIndent for readability)
	// TODO: Write to projects.json file
	panic("not implemented")
}

// AddProject adds a new project to the registry
// Returns error if project already exists
func AddProject(project string) error {
	// TODO: Call ReadProjects() to get existing projects
	// TODO: Check if project already exists (loop through slice)
	// TODO: If exists, return error: "project already exists"
	// TODO: Append new project to slice
	// TODO: Call WriteProjects() with updated slice
	panic("not implemented")
}

// RemoveProject removes a project from the registry
// Returns error if project doesn't exist
// Returns error if any tasks are using this project
func RemoveProject(project string) error {
	// TODO: Check if any tasks use this project
	//       - Call ReadTasks()
	//       - Loop through tasks and check their Project field
	//       - If any task has this project, return error: "cannot remove project: tasks still use it"
	// TODO: Call ReadProjects() to get existing projects
	// TODO: Create new slice without the project
	// TODO: If project wasn't found, return error: "project not found"
	// TODO: Call WriteProjects() with new slice
	panic("not implemented")
}

// ProjectExists checks if a project is registered
func ProjectExists(project string) (bool, error) {
	// TODO: Call ReadProjects()
	// TODO: Loop through projects to find match
	// TODO: Return true if found, false otherwise
	panic("not implemented")
}

// ValidateProject checks if a project is registered
// Returns error if project is not empty and not registered
// Empty project string is valid (means no project assigned)
func ValidateProject(project string) error {
	// TODO: If project is empty string, return nil (valid - no project)
	// TODO: Call ProjectExists()
	// TODO: If not exists, return error: "project 'xyz' not registered. Use: godo project add xyz"
	// TODO: Return nil if valid
	panic("not implemented")
}

// ListProjects returns all registered projects
// This is just an alias for ReadProjects for semantic clarity
func ListProjects() ([]string, error) {
	return ReadProjects()
}
