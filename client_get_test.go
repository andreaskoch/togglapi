package togglapi

import (
	"fmt"
	"io"
	"testing"
)

func Test_GetClients_RestClientReturnsError_ErrorIsReturned(t *testing.T) {
	// arrange
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return nil, fmt.Errorf("Some error")
		},
	}

	clientAPI := &ClientAPI{
		restClient: restClient,
	}

	// act
	_, err := clientAPI.GetClients()

	// assert
	if err == nil {
		t.Fail()
		t.Logf("GetClients should return an error if the rest client returned an error")
	}
}

func Test_GetClients_InvalidJSONIsReturned_ErrorIsReturned(t *testing.T) {
	// arrange
	clientsJSON := `
  {
    id: dasdasd

		dasdkals&/5758796`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(clientsJSON), nil
		},
	}

	clientAPI := &ClientAPI{
		restClient: restClient,
	}

	// act
	_, err := clientAPI.GetClients()

	// assert
	if err == nil {
		t.Fail()
		t.Logf("GetClients should return an error if the JSON returned by the API is invalid")
	}
}

func Test_GetClients_NoClientsReturned_EmptyListIsReturned(t *testing.T) {
	// arrange
	clientsJSON := `[]`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(clientsJSON), nil
		},
	}

	clientAPI := &ClientAPI{
		restClient: restClient,
	}

	// act
	clients, err := clientAPI.GetClients()

	// assert
	if len(clients) > 0 || err != nil {
		t.Fail()
		if err != nil {
			t.Logf("GetClients should not return an error if there are no clients but returned this: %s", err)
		} else {
			t.Logf("GetClients should return an empty list if the API returns no clients")
		}
	}
}

func Test_GetClients_HTTPMethodIsGET(t *testing.T) {
	// arrange
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			// assert
			if method != "GET" {
				t.Fail()
				t.Logf("GetClients should have used a GET request")
			}

			return nil, nil
		},
	}

	clientAPI := &ClientAPI{
		restClient: restClient,
	}

	// act
	clientAPI.GetClients()
}

func Test_GetClients_ValidJSONIsReturned_ClientsAreReturned(t *testing.T) {
	// arrange
	clientsJSON := `[
  {
    "id": 1,
    "wid": 1,
    "name": "Client A",
    "notes": ""
  },
	{
    "id": 2,
    "wid": 1,
    "name": "Client B",
    "notes": "Yada Yada"
  }
]`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(clientsJSON), nil
		},
	}

	clientAPI := &ClientAPI{
		restClient: restClient,
	}

	// act
	clients, err := clientAPI.GetClients()

	// assert
	if clients == nil || len(clients) != 2 {
		t.Fail()

		if err != nil {
			t.Logf("GetClients should have returned 2 clients but returned an error instead: %s", err)
		} else {
			t.Logf("GetClients should have returned 2 clients")
		}
	}
}
