package storage

// ReadTags reads all registered tags from tags.json
// File format: {"tags": ["backend", "frontend", "docs"]}
// Returns empty slice if file doesn't exist or is empty
func ReadTags() ([]string, error) {
	// TODO: Get workspace path
	// TODO: Read tags.json file
	// TODO: Unmarshal into TagRegistry struct
	// TODO: Return the Tags slice
	// TODO: If file doesn't exist, return empty slice (not an error)
	panic("not implemented")
}

// WriteTags overwrites tags.json with the provided tags
func WriteTags(tags []string) error {
	// TODO: Get workspace path
	// TODO: Create TagRegistry struct with tags
	// TODO: Marshal to JSON with indentation (json.MarshalIndent for readability)
	// TODO: Write to tags.json file
	panic("not implemented")
}

// AddTag adds a new tag to the registry
// Returns error if tag already exists
func AddTag(tag string) error {
	// TODO: Call ReadTags() to get existing tags
	// TODO: Check if tag already exists (loop through slice)
	// TODO: If exists, return error: "tag already exists"
	// TODO: Append new tag to slice
	// TODO: Call WriteTags() with updated slice
	panic("not implemented")
}

// RemoveTag removes a tag from the registry
// Returns error if tag doesn't exist
// Returns error if any tasks are using this tag
func RemoveTag(tag string) error {
	// TODO: Check if any tasks use this tag
	//       - Call ReadTasks()
	//       - Loop through tasks and check their Tags slice
	//       - If any task has this tag, return error: "cannot remove tag: tasks still use it"
	// TODO: Call ReadTags() to get existing tags
	// TODO: Create new slice without the tag
	// TODO: If tag wasn't found, return error: "tag not found"
	// TODO: Call WriteTags() with new slice
	panic("not implemented")
}

// TagExists checks if a tag is registered
func TagExists(tag string) (bool, error) {
	// TODO: Call ReadTags()
	// TODO: Loop through tags to find match
	// TODO: Return true if found, false otherwise
	panic("not implemented")
}

// ValidateTags checks if all provided tags are registered
// Returns error with details if any tag is not registered
func ValidateTags(tags []string) error {
	// TODO: Call ReadTags() to get registered tags
	// TODO: For each provided tag:
	//       - Check if it exists in registered tags
	//       - If not, return error: "tag 'xyz' not registered. Use: godo tag add xyz"
	// TODO: Return nil if all tags are valid
	panic("not implemented")
}

// ListTags returns all registered tags
// This is just an alias for ReadTags for semantic clarity
func ListTags() ([]string, error) {
	return ReadTags()
}
