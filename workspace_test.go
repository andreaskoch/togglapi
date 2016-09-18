package togglapi

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_NewWorkspaceRepository(t *testing.T) {
	// act
	client := NewWorkspaceRepository("http://api.example.com", "sakldjaksljkl312312")

	// assert
	if client == nil {
		t.Fail()
		t.Logf("NewWorkspaceRepository should have returned a workspace API client")
	}
}

// If you are only interested in the Workspace API you can instantiate a
// WorkspaceRepository using the NewWorkspaceRepository function.
func ExampleNewWorkspaceRepository() {
	apiToken := "Your-Toggl-API-Token"
	baseURL := "https://www.toggl.com/api/v8"
	workspaceRepository := NewWorkspaceRepository(baseURL, apiToken)

	workspaces, workspacesError := workspaceRepository.GetWorkspaces()
	if workspacesError != nil {
		fmt.Fprintf(os.Stderr, "Failed to get workspaces: %s", workspacesError)
		return
	}

	for _, workspace := range workspaces {
		fmt.Println(workspace.Name)
	}
}

func Test_GetWorkspaces_RestClientReturnsError_ErrorIsReturned(t *testing.T) {
	// arrange
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return nil, fmt.Errorf("Some error")
		},
	}

	repository := &RESTWorkspaceRepository{
		restClient: restClient,
	}

	// act
	_, err := repository.GetWorkspaces()

	// assert
	if err == nil {
		t.Fail()
		t.Logf("GetWorkspaces should return an error if the rest client returns one")
	}
}

func Test_GetWorkspaces_RestClientReturnsInvalidJSON_ErrorIsReturned(t *testing.T) {
	// arrange
	workspacesJSON := `[
  {;,,,.,daskdlasdlak ---invalid--
    "id": 1,
    "name": "TogglCSV's workspace"
  }
]`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(workspacesJSON), nil
		},
	}

	repository := &RESTWorkspaceRepository{
		restClient: restClient,
	}

	// act
	_, err := repository.GetWorkspaces()

	// assert
	if err == nil || !strings.Contains(err.Error(), "Failed to deserialize") {
		t.Fail()
		t.Logf("GetWorkspaces should return an error if the JSON returned by the API is invalid")
	}
}

func Test_GetWorkspaces_NoWorkspacesReturned_EmptyListIsReturned(t *testing.T) {
	// arrange
	workspacesJSON := `[]`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(workspacesJSON), nil
		},
	}

	repository := &RESTWorkspaceRepository{
		restClient: restClient,
	}

	// act
	workspaces, err := repository.GetWorkspaces()

	// assert
	if len(workspaces) > 0 || err != nil {
		t.Fail()

		if err != nil {
			t.Logf("GetWorkspaces should not return an error if there are no workspaces but returned this: %s", err)
		} else {
			t.Logf("GetWorkspaces should return an empty list if the API returns no workspaces")
		}

	}
}

func Test_GetWorkspaces_HTTPMethodIsGET(t *testing.T) {
	// arrange
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			// assert
			if method != "GET" {
				t.Fail()
				t.Logf("GetWorkspaces should issue a GET request")
			}

			return []byte(""), nil
		},
	}

	repository := &RESTWorkspaceRepository{
		restClient: restClient,
	}

	// act
	repository.GetWorkspaces()
}

func Test_GetWorkspaces_ValidJSONIsReturned_ProjectsAreReturned(t *testing.T) {
	// arrange
	workspacesJSON := `[
  {
    "id": 1,
    "name": "TogglCSV's workspace"
  },
	{
    "id": 2,
    "name": "TogglCSV's second workspace"
  }
]`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(workspacesJSON), nil
		},
	}

	repository := &RESTWorkspaceRepository{
		restClient: restClient,
	}

	// act
	workspaces, err := repository.GetWorkspaces()

	// assert
	if workspaces == nil || len(workspaces) != 2 {
		t.Fail()

		if err != nil {
			t.Logf("GetWorkspaces should have returned 2 workspaces but returned an error instead: %s", err)
		} else {
			t.Logf("GetWorkspaces should have returned 2 workspaces")
		}
	}
}
