package togglapi

import (
	"fmt"
	"os"
	"testing"
)

func Test_NewProjectRepository(t *testing.T) {
	// act
	repository := NewProjectRepository("http://api.example.com", "sakldjaksljkl312312")

	// assert
	if repository == nil {
		t.Fail()
		t.Logf("NewProjectRepository should have returned a project API client")
	}
}

// If you are only interested in the Project API you can instantiate a
// ProjectRepository using the NewProjectRepository function.
func ExampleNewProjectRepository() {
	apiToken := "Your-Toggl-API-Token"
	baseURL := "https://www.toggl.com/api/v8"
	projectRepository := NewProjectRepository(baseURL, apiToken)

	workspaces, workspacesError := projectRepository.GetWorkspaces()
	if workspacesError != nil {
		fmt.Fprintf(os.Stderr, "Failed to get workspaces: %s", workspacesError)
		return
	}

	for _, workspace := range workspaces {

		projects, projectsError := api.GetProjects(workspace.ID)
		if projectsError != nil {
			fmt.Fprintf(os.Stderr, "Failed to get projects: %s", projectsError)
			return
		}

		for _, project := range projects {
			fmt.Println(project.Name)
		}
	}
}
