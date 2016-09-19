package togglapi

import (
	"fmt"
	"io"
	"testing"
)

func Test_GetProjects_RestClientReturnsError_ErrorIsReturned(t *testing.T) {
	// arrange
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return nil, fmt.Errorf("Some error")
		},
	}

	projectAPI := &ProjectAPI{
		restClient: restClient,
	}

	// act
	_, err := projectAPI.GetProjects(1)

	// assert
	if err == nil {
		t.Fail()
		t.Logf("GetProjects should return an error if the rest client returned an error")
	}
}

func Test_GetProjects_InvalidJSONIsReturned_ErrorIsReturned(t *testing.T) {
	// arrange
	projectsJSON := `
  {
    id: 22126959
    wid: 1641370
    name: Meetings
    billable: false
    is_private: true
    active: true
    template: false
    at: 2016-09-06T09:32:06+00:00
    created_at: 2016-09-06T09:32:06+00:00
    color: 2
    auto_estimates: false
    actual_hours: 338
    hex_color: #df7baa
  }`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(projectsJSON), nil
		},
	}

	projectAPI := &ProjectAPI{
		restClient: restClient,
	}

	// act
	_, err := projectAPI.GetProjects(1)

	// assert
	if err == nil {
		t.Fail()
		t.Logf("GetProjects should return an error if the JSON returned by the API is invalid")
	}
}

func Test_GetProjects_NoProjectsReturned_EmptyListIsReturned(t *testing.T) {
	// arrange
	projectsJSON := `[]`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(projectsJSON), nil
		},
	}

	projectAPI := &ProjectAPI{
		restClient: restClient,
	}

	// act
	projects, err := projectAPI.GetProjects(1)

	// assert
	if len(projects) > 0 || err != nil {
		t.Fail()
		if err != nil {
			t.Logf("GetProjects should not return an error if there are no projects but returned this: %s", err)
		} else {
			t.Logf("GetProjects should return an empty list if the API returns no projects")
		}
	}
}

func Test_GetProjects_HTTPMethodIsGET(t *testing.T) {
	// arrange
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			// assert
			if method != "GET" {
				t.Fail()
				t.Logf("GetProjects should have used a GET request")
			}

			return nil, nil
		},
	}

	projectAPI := &ProjectAPI{
		restClient: restClient,
	}

	// act
	projectAPI.GetProjects(1)
}

func Test_GetProjects_ValidJSONIsReturned_ProjectsAreReturned(t *testing.T) {
	// arrange
	projectsJSON := `[
  {
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
  },
  {
    "id": 2,
    "wid": 1,
    "name": "Bugs",
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
  }
]`
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(projectsJSON), nil
		},
	}

	projectAPI := &ProjectAPI{
		restClient: restClient,
	}

	// act
	projects, err := projectAPI.GetProjects(1)

	// assert
	if projects == nil || len(projects) != 2 {
		t.Fail()

		if err != nil {
			t.Logf("GetProjects should have returned 2 projects but returned an error instead: %s", err)
		} else {
			t.Logf("GetProjects should have returned 2 projects")
		}
	}
}
