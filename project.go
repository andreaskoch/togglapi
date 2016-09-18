package togglapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andreaskoch/togglapi/model"
	"github.com/pkg/errors"
)

// NewProjectRepository create a new client for the Toggl project API.
func NewProjectRepository(baseURL, token string) model.ProjectRepository {
	return &RESTProjectRepository{
		restClient: &togglRESTAPIClient{
			baseURL: baseURL,
			token:   token,
		},
	}
}

// RESTProjectRepository provides functions for interacting with Toggls' project API.
type RESTProjectRepository struct {
	restClient RESTRequester
}

// CreateProject creates a new project.
func (repository *RESTProjectRepository) CreateProject(project model.Project) (model.Project, error) {

	projectRequest := struct {
		Project model.Project `json:"project"`
	}{
		Project: project,
	}

	jsonBody, marshalError := json.Marshal(projectRequest)
	if marshalError != nil {
		return model.Project{}, errors.Wrap(marshalError, "Failed to serialize the project")
	}

	content, err := repository.restClient.Request(http.MethodPost, "projects", bytes.NewBuffer(jsonBody))
	if err != nil {
		return model.Project{}, errors.Wrap(err, "Failed to create project")
	}

	var createdProject model.Project
	if unmarshalError := json.Unmarshal(content, &createdProject); unmarshalError != nil {
		return model.Project{}, errors.Wrap(unmarshalError, "Failed to deserialize the created project")
	}

	return createdProject, nil
}

// GetProjects returns all projects for the given workspace.
func (repository *RESTProjectRepository) GetProjects(workspaceID int) ([]model.Project, error) {

	route := fmt.Sprintf(
		"workspaces/%d/projects",
		workspaceID,
	)

	content, err := repository.restClient.Request(http.MethodGet, route, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve projects")
	}

	var projects []model.Project
	if unmarshalError := json.Unmarshal(content, &projects); unmarshalError != nil {
		return nil, errors.Wrap(unmarshalError, "Failed to deserialize the projects")
	}

	return projects, nil
}
