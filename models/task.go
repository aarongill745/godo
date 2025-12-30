package models

import "time"

// Task represents an active todo task
type Task struct {
	ID          int       `json:"id"`
	Project     string    `json:"project"` // Empty string if no project
	Description string    `json:"desc"`
	Priority    int       `json:"priority"` // Count of '+' symbols (0 = lowest)
	Created     time.Time `json:"created"`
	Tags        []string  `json:"tags"` // Registered tags only
}

// CompletedTask represents a completed task with timestamp
type CompletedTask struct {
	Task
	CompletedAt time.Time `json:"completed_at"`
}

// TagRegistry holds registered tag names
type TagRegistry struct {
	Tags []string `json:"tags"`
}

// ProjectRegistry holds registered project names
type ProjectRegistry struct {
	Projects []string `json:"projects"`
}
