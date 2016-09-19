package togglapi

import (
	"fmt"
	"os"
	"testing"
)

func Test_NewProjectAPI(t *testing.T) {
	// act
	projectAPI := NewProjectAPI("http://api.example.com", "sakldjaksljkl312312")

	// assert
	if projectAPI == nil {
		t.Fail()
		t.Logf("NewProjectAPI should have returned a project API client")
	}
}

// If you are only interested in the Project API you can instantiate a
// ProjectAPI using the NewProjectAPI function.
func ExampleNewProjectAPI() {
	apiToken := "Your-Toggl-API-Token"
	baseURL := "https://www.toggl.com/api/v8"

	workspaceAPI := NewWorkspaceAPI(baseURL, apiToken)
	workspaces, workspacesError := workspaceAPI.GetWorkspaces()
	if workspacesError != nil {
		fmt.Fprintf(os.Stderr, "Failed to get workspaces: %s", workspacesError)
		return
	}

	projectRepository := NewProjectAPI(baseURL, apiToken)
	for _, workspace := range workspaces {

		projects, projectsError := projectRepository.GetProjects(workspace.ID)
		if projectsError != nil {
			fmt.Fprintf(os.Stderr, "Failed to get projects: %s", projectsError)
			return
		}

		for _, project := range projects {
			fmt.Println(project.Name)
		}
	}
}
