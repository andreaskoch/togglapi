package togglapi

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// The RESTRequester interface provides a function for sending HTTP requests
// to REST APIs.
type RESTRequester interface {
	// Request sends an HTTP request with the given parameters (method, route, payload)
	// to an REST API and returns the APIs' response or an error if the request failed.
	Request(method, route string, payload io.Reader) ([]byte, error)
}

// The togglRESTAPIClient perform the HTTP requests against the Toggl API and
// returns the APIs' response.
type togglRESTAPIClient struct {
	baseURL string
	token   string
}

// Request sends an HTTP request with the given parameters (method, route, payload) to the Toggl
// API and returns the APIs' response or an error if the request failed.
func (client *togglRESTAPIClient) Request(method, route string, payload io.Reader) ([]byte, error) {

	httpClient := &http.Client{}

	actionURL := fmt.Sprintf(
		"%s/%s",
		client.baseURL,
		route,
	)

	req, err := http.NewRequest(method, actionURL, payload)
	if err != nil {
		return nil, err
	}

	// add the API token
	req.SetBasicAuth(client.token, "api_token")

	// execute the request
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	content, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		return nil, errors.Wrap(readError, "Failed to read response body")
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("The %s request against %s failed (%s): %s", req.Method, req.URL, response.Status, content)
	}

	return content, nil

}
