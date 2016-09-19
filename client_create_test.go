package togglapi

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/andreaskoch/togglapi/model"
)

func Test_CreateClient_RestClientReturnsError_ErrorIsReturned(t *testing.T) {
	// arrange
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return nil, fmt.Errorf("Some error")
		},
	}

	clientAPI := &ClientAPI{
		restClient: restClient,
	}

	input := model.Client{}

	// act
	_, err := clientAPI.CreateClient(input)

	// assert
	if err == nil {
		t.Fail()
		t.Logf("CreateClient should return an error if the rest client returns an error")
	}
}

func Test_CreateClient_InvalidJSONIsReturned_ErrorIsReturned(t *testing.T) {
	// arrange
	clientsJSON := `dsakdlajkl,,d;; jkjk??`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(clientsJSON), nil
		},
	}

	clientAPI := &ClientAPI{
		restClient: restClient,
	}

	input := model.Client{}

	// act
	_, err := clientAPI.CreateClient(input)

	// assert
	if err == nil || !strings.Contains(err.Error(), "Failed to deserialize the created client") {
		t.Fail()
		t.Logf("CreateClient should return an error if the JSON returned by the API is invalid")
	}
}

func Test_CreateClient_NoClientsReturned_EmptyClientIsReturned(t *testing.T) {
	// arrange
	clientJSON := `null`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(clientJSON), nil
		},
	}

	clientAPI := &ClientAPI{
		restClient: restClient,
	}

	input := model.Client{}

	// act
	_, err := clientAPI.CreateClient(input)

	// assert
	if err != nil {
		t.Fail()
		t.Logf("CreateClient should not return an error: %s", err)
	}
}

func Test_CreateClient_HTTPMethodIsPOST(t *testing.T) {
	// arrange
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			// assert
			if method != "POST" {
				t.Fail()
				t.Logf("CreateClient should issue a POST request")
			}

			return nil, nil
		},
	}

	clientAPI := &ClientAPI{
		restClient: restClient,
	}

	input := model.Client{}

	// act
	clientAPI.CreateClient(input)
}

func Test_CreateClient_ValidJSONIsReturned_CreatedClientIsReturned(t *testing.T) {
	// arrange
	clientJSON := `{
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
			return []byte(clientJSON), nil
		},
	}

	clientAPI := &ClientAPI{
		restClient: restClient,
	}

	input := model.Client{}

	// act
	client, err := clientAPI.CreateClient(input)

	// assert
	if err != nil || client.Name != "Meetings" {
		t.Fail()

		if err != nil {
			t.Logf("CreateClient should have returned a client but returned an error instead: %s", err)
		} else {
			t.Logf("CreateClient should have returned a client but returned this: %#v", client)
		}
	}
}
