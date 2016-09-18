package togglapi

import (
	"encoding/json"
	"net/http"

	"github.com/andreaskoch/togglapi/model"
	"github.com/pkg/errors"
)

// NewWorkspaceRepository create a new client for the Toggl workspace API.
func NewWorkspaceRepository(baseURL, token string) model.WorkspaceRepository {
	return &RESTWorkspaceRepository{
		restClient: &togglRESTAPIClient{
			baseURL: baseURL,
			token:   token,
		},
	}
}

// RESTWorkspaceRepository provides functions for interacting with Toggls' workspace API.
type RESTWorkspaceRepository struct {
	restClient RESTRequester
}

// GetWorkspaces returns all workspaces for the current user.
func (repository *RESTWorkspaceRepository) GetWorkspaces() ([]model.Workspace, error) {
	content, err := repository.restClient.Request(http.MethodGet, "workspaces", nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve workspaces")
	}

	var workspaces []model.Workspace
	if unmarshalError := json.Unmarshal(content, &workspaces); unmarshalError != nil {
		return nil, errors.Wrap(unmarshalError, "Failed to deserialize the workspaces")
	}

	return workspaces, nil
}
