// Package togglapi provides repsitorties for accessing Toggl objects
// such as Workspaces, Projects and Time Entries via the Toggl REST API.
package togglapi

import (
	"github.com/andreaskoch/togglapi/date"
	"github.com/andreaskoch/togglapi/model"
)

const clientName = "github.com/andreaskoch/togglapi"

// NewAPI create a new instance of the Toggl API.
func NewAPI(baseURL, token string) model.TogglAPI {
	restAPI := &togglRESTAPIClient{
		baseURL: baseURL,
		token:   token,
	}

	dateFormatter := date.NewISO8601Formatter()

	return &API{
		&RESTWorkspaceRepository{restAPI},
		&RESTProjectRepository{restAPI},
		&RESTTimeEntryRepository{restAPI, dateFormatter},
		&RESTClientRepository{restAPI},
	}
}

// API provides functions for interacting with the Toggl API.
type API struct {
	model.WorkspaceRepository
	model.ProjectRepository
	model.TimeEntryRepository
	model.ClientRepository
}
