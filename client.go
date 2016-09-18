package togglapi

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/andreaskoch/togglapi/model"
	"github.com/pkg/errors"
)

// NewClientRepository create a new client for the Toggl client API.
func NewClientRepository(baseURL, token string) model.ClientRepository {
	return &RESTClientRepository{
		restClient: &togglRESTAPIClient{
			baseURL: baseURL,
			token:   token,
		},
	}
}

// RESTClientRepository provides functions for interacting with Toggls' client API.
type RESTClientRepository struct {
	restClient RESTRequester
}

// CreateClient creates a new client.
func (repository *RESTClientRepository) CreateClient(client model.Client) (model.Client, error) {

	clientRequest := struct {
		Client model.Client `json:"client"`
	}{
		Client: client,
	}

	jsonBody, marshalError := json.Marshal(clientRequest)
	if marshalError != nil {
		return model.Client{}, errors.Wrap(marshalError, "Failed to serialize the client")
	}

	content, err := repository.restClient.Request(http.MethodPost, "clients", bytes.NewBuffer(jsonBody))
	if err != nil {
		return model.Client{}, errors.Wrap(err, "Failed to create client")
	}

	var createdClient model.Client
	if unmarshalError := json.Unmarshal(content, &createdClient); unmarshalError != nil {
		return model.Client{}, errors.Wrap(unmarshalError, "Failed to deserialize the created client")
	}

	return createdClient, nil
}

// GetClients returns all clients for the given workspace.
func (repository *RESTClientRepository) GetClients() ([]model.Client, error) {
	content, err := repository.restClient.Request(http.MethodGet, "clients", nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve clients")
	}

	var clients []model.Client
	if unmarshalError := json.Unmarshal(content, &clients); unmarshalError != nil {
		return nil, errors.Wrap(unmarshalError, "Failed to deserialize the clients")
	}

	return clients, nil
}
