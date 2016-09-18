package togglapi

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/andreaskoch/togglapi/model"
)

func Test_CreateProject_RestClientReturnsError_ErrorIsReturned(t *testing.T) {
	// arrange
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return nil, fmt.Errorf("Some error")
		},
	}

	repository := &RESTProjectRepository{
		restClient: restClient,
	}

	input := model.Project{}

	// act
	_, err := repository.CreateProject(input)

	// assert
	if err == nil {
		t.Fail()
		t.Logf("CreateProject should return an error if the rest client returns an error")
	}
}

func Test_CreateProject_InvalidJSONIsReturned_ErrorIsReturned(t *testing.T) {
	// arrange
	projectsJSON := `dsakdlajkl,,d;; jkjk??`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(projectsJSON), nil
		},
	}

	repository := &RESTProjectRepository{
		restClient: restClient,
	}

	input := model.Project{}

	// act
	_, err := repository.CreateProject(input)

	// assert
	if err == nil || !strings.Contains(err.Error(), "Failed to deserialize the created project") {
		t.Fail()
		t.Logf("CreateProject should return an error if the JSON returned by the API is invalid")
	}
}

func Test_CreateProject_NoProjectsReturned_EmptyProjectIsReturned(t *testing.T) {
	// arrange
	projectJSON := `null`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(projectJSON), nil
		},
	}

	repository := &RESTProjectRepository{
		restClient: restClient,
	}

	input := model.Project{}

	// act
	_, err := repository.CreateProject(input)

	// assert
	if err != nil {
		t.Fail()
		t.Logf("CreateProject should not return an error: %s", err)
	}
}

func Test_CreateProject_HTTPMethodIsPOST(t *testing.T) {
	// arrange
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			// assert
			if method != "POST" {
				t.Fail()
				t.Logf("CreateProject should issue a POST request")
			}

			return nil, nil
		},
	}

	repository := &RESTProjectRepository{
		restClient: restClient,
	}

	input := model.Project{}

	// act
	repository.CreateProject(input)
}

func Test_CreateProject_ValidJSONIsReturned_CreatedProjectIsReturned(t *testing.T) {
	// arrange
	projectJSON := `{
  "id": 1,
  "wid": 1,
  "name": "Meetings",
  "billable": false,
  "is_private": true,
  "active": true,
  "template": false,
  "at": "2016-09-06T09:32:06+00:00",
  "created_at": "2016-09-06T09:32:06+00:00",
  "color": "2",
  "auto_estimates": false,
  "actual_hours": 338,
  "hex_color": "#df7baa"
}`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(projectJSON), nil
		},
	}

	repository := &RESTProjectRepository{
		restClient: restClient,
	}

	input := model.Project{}

	// act
	project, err := repository.CreateProject(input)

	// assert
	if err != nil || project.Name != "Meetings" {
		t.Fail()

		if err != nil {
			t.Logf("CreateProject should have returned a project but returned an error instead: %s", err)
		} else {
			t.Logf("CreateProject should have returned a project but returned this: %#v", project)
		}
	}
}
