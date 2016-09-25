package togglapi

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type mockRESTRequester struct {
	request func(method, route string, payload io.Reader) ([]byte, error)
}

func (requester *mockRESTRequester) Request(method, route string, payload io.Reader) ([]byte, error) {
	return requester.request(method, route, payload)
}

func Test_Request_APIReturns404Error_ErrorIsReturned(t *testing.T) {
	// arrange
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))

	defer testServer.Close()

	restClient := &togglRESTAPIClient{
		baseURL: testServer.URL,
		token:   "21das6d567a5d67s",
	}

	// assert
	_, err := restClient.Request("GET", "some-route", nil)

	// assert
	if err == nil {
		t.Fail()
		t.Logf("Request should have returned an error if the HTTP API returned an 404 error")
	}
}

func Test_Request_APIReturns500Error_ErrorIsReturned(t *testing.T) {
	// arrange
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "500 Internal Server Error", 500)
	}))

	defer testServer.Close()

	restClient := &togglRESTAPIClient{
		baseURL: testServer.URL,
		token:   "21das6d567a5d67s",
	}

	// assert
	_, err := restClient.Request("GET", "some-route", nil)

	// assert
	if err == nil {
		t.Fail()
		t.Logf("Request should have returned an error if the HTTP API returned an 500 error")
	}
}

func Test_Request_ResponseIsNil_ContentIsEmpty(t *testing.T) {
	// arrange
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	defer testServer.Close()

	restClient := &togglRESTAPIClient{
		baseURL: testServer.URL,
		token:   "21das6d567a5d67s",
	}

	// assert
	content, _ := restClient.Request("GET", "some-route", nil)

	// assert
	if len(content) > 0 {
		t.Fail()
		t.Logf("Request should not have returned a response")
	}
}

func Test_Request_CalledMultipleTimes_RequestRateLimitApplies(t *testing.T) {
	// arrange
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	defer testServer.Close()

	restClient := &togglRESTAPIClient{
		baseURL:              testServer.URL,
		token:                "21das6d567a5d67s",
		pauseBetweenRequests: time.Millisecond * 33,
	}

	// assert
	timeOfFirstRequest := time.Now()

	restClient.Request("GET", "some-route", nil)
	restClient.Request("GET", "some-route", nil)
	restClient.Request("GET", "some-route", nil)

	// assert
	duration := time.Since(timeOfFirstRequest)
	expectedDuration := time.Millisecond * 33 * 2
	if duration < expectedDuration {
		t.Fail()
		t.Logf("Issueing 3 requests should have taken at least 66ms but took only %s", duration)
	}
}
