# Godo Rewrite - Implementation Plan

This document tracks the incremental implementation of the godo rewrite. We're building features one at a time to keep commits clean and learning focused.

## Project Structure (Current)

```
godo/
├── main.go           # Cobra root command + init subcommand
├── workspace.go      # Workspace discovery functions
├── go.mod
├── go.sum
└── old/              # Previous implementation (preserved)
```

## Implementation Phases

### Phase 1: Foundation ✓ (In Progress)

**Goal:** Get workspace discovery and initialization working

- [x] Install cobra dependency
- [x] Create main.go with cobra setup
- [ ] Implement workspace.go with FindWorkspace() and WorkspaceExists()
- [ ] Test `godo init` creates .godo/
- [ ] Test workspace discovery from subdirectories

**What we're learning:**
- Cobra command structure
- Directory traversal in Go
- File/directory operations

---

### Phase 2: Tag System

**Goal:** Tag registry management (prerequisite for adding tasks with tags)

**Files to create:**
- `tags.go` - Tag storage operations (read/write tags.json)
- Update `main.go` - Add tag subcommands

**Commands to implement:**
```bash
godo tag add <name>       # Register new tag
godo tag list             # Show all registered tags
godo tag remove <name>    # Remove tag (check if in use)
```

**Data structure:**
```
.godo/tags.json:
["backend", "frontend", "docs"]
```

**What we'll learn:**
- JSON marshaling/unmarshaling
- Slice operations in Go
- Command validation logic

---

### Phase 3: Task Model & Storage

**Goal:** Define task structure and JSONL operations

**Files to create:**
- `task.go` - Task and CompletedTask structs
- `storage.go` - JSONL read/write operations

**Task structure:**
```go
type Task struct {
    ID          int       `json:"id"`
    Description string    `json:"desc"`
    Priority    int       `json:"priority"`
    Created     time.Time `json:"created"`
    Tags        []string  `json:"tags"`
}
```

**Functions needed:**
- `LoadTasks(workspacePath string) ([]Task, error)`
- `SaveTasks(workspacePath string, tasks []Task) error`
- `AppendTask(workspacePath string, task Task) error`
- `GetNextTaskID(workspacePath string) (int, error)`

**What we'll learn:**
- JSONL format (one JSON object per line)
- Time handling in Go
- Error handling patterns

---

### Phase 4: Add Command

**Goal:** Create tasks with priority and tags

**Files to update:**
- `main.go` - Add `add` command

**Command syntax:**
```bash
godo add "description" [+[+[+]]] [@tag1 @tag2...]
godo add "Fix bug" +++ @backend @urgent
```

**Implementation steps:**
1. Parse command arguments (description, priority, tags)
2. Validate tags exist in registry
3. Create Task struct with generated ID
4. Append to tasks.jsonl

**What we'll learn:**
- Argument parsing
- String manipulation (extracting +, @symbols)
- Tag validation logic

---

### Phase 5: List Command (Basic)

**Goal:** Display all tasks without filtering

**Files to update:**
- `main.go` - Add `list` command

**Command syntax:**
```bash
godo list    # Show all tasks sorted by ID
```

**Display format:**
```
ID  Priority  Description              Tags
1   +++       Fix login bug            @backend @urgent
2   +         Update documentation     @docs
3             Refactor tests           @backend
```

**What we'll learn:**
- Text formatting and alignment
- ANSI colors for terminal output
- Table rendering

---

### Phase 6: List Command (Filtering & Sorting)

**Goal:** Add filtering and sorting capabilities

**Files to update:**
- `main.go` - Add flags to list command

**Command syntax:**
```bash
godo list --tag backend           # Filter by tag
godo list --priority 3            # Filter by priority
godo list --sort priority         # Sort by priority (desc)
godo list --sort created          # Sort by creation date
```

**What we'll learn:**
- Cobra flag binding
- Slice filtering
- Custom sorting with sort.Slice()

---

### Phase 7: Complete Command

**Goal:** Mark tasks as complete and move to history

**Files to create:**
- Update `storage.go` - Add completed task operations

**Command syntax:**
```bash
godo complete <id>
```

**Implementation:**
1. Find task by ID in tasks.jsonl
2. Create CompletedTask with timestamp
3. Append to completed.jsonl
4. Remove from tasks.jsonl (rewrite file without that task)

**What we'll learn:**
- File rewriting strategies
- Timestamp handling
- Task ID lookups

---

### Phase 8: History Command

**Goal:** View completed tasks

**Files to update:**
- `main.go` - Add `history` command

**Command syntax:**
```bash
godo history                # All completed tasks
godo history --tag backend  # Completed tasks by tag
```

**What we'll learn:**
- Reusing list display logic
- Working with completed tasks

---

### Phase 9: Edit Command

**Goal:** Edit task descriptions

**Files to update:**
- `main.go` - Add `edit` command

**Command syntax:**
```bash
godo edit <id> "new description"
```

**Implementation:**
1. Load all tasks
2. Find task by ID
3. Update description
4. Save all tasks back to file
5. Show before/after confirmation

**What we'll learn:**
- In-place updates in JSONL
- Displaying diffs

---

### Phase 10: Profile System

**Goal:** Quick initialization with predefined tag sets

**Files to update:**
- `main.go` - Add `--profile` flag to init command

**Command syntax:**
```bash
godo init --profile developer
godo init --profile personal
godo init --profile research
```

**Profiles:**
- **developer:** @frontend, @backend, @infra, @cicd, @refactor, @testing, @docs, @bugfix
- **personal:** @home, @work, @errands, @shopping, @health, @finance
- **research:** @reading, @writing, @experiments, @analysis, @review

**What we'll learn:**
- Map data structures for profiles
- Combining init + tag creation

---

### Phase 11: Help System

**Goal:** Comprehensive help documentation

**Files to update:**
- `main.go` - Enhance help text for all commands

**Command syntax:**
```bash
godo help
godo help add
godo help list
```

**What we'll learn:**
- Cobra's help system
- Documentation best practices

---

### Phase 12: Installation

**Goal:** Global installation support

**Files to create:**
- `Makefile` or install script

**Installation methods:**
```bash
make install              # Copy to /usr/local/bin
go install               # For Go users
```

**What we'll learn:**
- Build process
- Installation paths
- Distribution methods

---

## Testing Strategy

For each phase:
1. Build: `go build`
2. Manual testing of the new feature
3. Test edge cases (empty workspace, invalid input, etc.)
4. Test integration with previous features

Future: Add automated tests

---

## Current Status

**Phase:** 1 (Foundation)
**Next Step:** Complete workspace.go implementation
**Blocked:** None

---

## Notes & Decisions

- **File structure:** Starting flat (package main), will refactor to packages if needed
- **ID strategy:** Auto-increment based on max ID in tasks.jsonl
- **JSONL format:** One task per line for easy appending
- **Tag validation:** Hard fail with helpful error message
- **No nested workspaces:** Always use first .godo found walking up