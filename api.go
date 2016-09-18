// Package togglapi provides access to Toggls' time tracking API.
// The togglapi package provides functions for creating and retrieving
// workspaces, clients, projects and time entries.
//
// To learn more about the Toggl API visit:
// https://github.com/toggl/toggl_api_docs
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
