package model

import "time"

// The ProjectRepository interface provides functions for creating and fetching projects.
type ProjectRepository interface {
	// CreateProject creates a new project.
	CreateProject(project Project) (Project, error)

	// GetProjects returns all projects for the given workspace.
	GetProjects(workspaceID int) ([]Project, error)
}

// The ClientRepository interface provides functions for creating and fetching clients.
type ClientRepository interface {
	// CreateClient creates a new client.
	CreateClient(client Client) (Client, error)

	// GetClients returns all clients.
	GetClients() ([]Client, error)
}

// The WorkspaceRepository interface provides functions for fetching workspacs.
type WorkspaceRepository interface {
	// GetWorkspaces returns all workspaces for the current user.
	GetWorkspaces() ([]Workspace, error)
}

// The TimeEntryRepository interface provides functions for fetching and creating time entries.
type TimeEntryRepository interface {
	// CreateTimeEntry creates a new time entry.
	CreateTimeEntry(timeEntry TimeEntry) (TimeEntry, error)

	// GetTimeEntries returns all time entries created between the given start and end date.
	// Returns nil and an error if the time entries could not be retrieved.
	GetTimeEntries(start, end time.Time) ([]TimeEntry, error)
}

// A TogglAPI interface implements some of the Toggl API methods.
type TogglAPI interface {
	WorkspaceRepository
	ProjectRepository
	TimeEntryRepository
	ClientRepository
}
