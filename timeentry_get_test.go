package togglapi

import (
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/andreaskoch/togglapi/date"
)

func Test_GetTimeEntries_RestClientReturnsError_ErrorIsReturned(t *testing.T) {
	// arrange
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return nil, fmt.Errorf("Some error")
		},
	}

	timeEntryAPI := &TimeEntryAPI{
		restClient:    restClient,
		dateFormatter: date.NewISO8601Formatter(),
	}

	start := time.Date(2016, 1, 1, 0, 0, 1, 0, time.UTC)
	end := time.Date(2016, 6, 30, 23, 59, 59, 0, time.UTC)

	// act
	_, err := timeEntryAPI.GetTimeEntries(start, end)

	// assert
	if err == nil {
		t.Fail()
		t.Logf("GetTimeEntries should return an error if the REST client returned an error")
	}
}

func Test_GetTimeEntries_InvalidJSONIsReturned_ErrorIsReturned(t *testing.T) {
	// arrange
	timeEntriesJSON := `[
  {;,,,.,daskdlasdlak ---invalid--
    "id": 1,
]`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(timeEntriesJSON), nil
		},
	}

	timeEntryAPI := &TimeEntryAPI{
		restClient:    restClient,
		dateFormatter: date.NewISO8601Formatter(),
	}

	start := time.Date(2016, 1, 1, 0, 0, 1, 0, time.UTC)
	end := time.Date(2016, 6, 30, 23, 59, 59, 0, time.UTC)

	// act
	_, err := timeEntryAPI.GetTimeEntries(start, end)

	// assert
	if err == nil {
		t.Fail()
		t.Logf("GetTimeEntries should return an error if the JSON returned by the API is invalid")
	}
}

func Test_GetTimeEntries_NoTimeEntriesReturned_EmptyListIsReturned(t *testing.T) {
	// arrange
	timeEntriesJSON := `[]`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(timeEntriesJSON), nil
		},
	}

	timeEntryAPI := &TimeEntryAPI{
		restClient:    restClient,
		dateFormatter: date.NewISO8601Formatter(),
	}

	start := time.Date(2016, 1, 1, 0, 0, 1, 0, time.UTC)
	end := time.Date(2016, 6, 30, 23, 59, 59, 0, time.UTC)

	// act
	timeEntries, err := timeEntryAPI.GetTimeEntries(start, end)

	// assert
	if len(timeEntries) > 0 || err != nil {
		t.Fail()

		if err != nil {
			t.Logf("GetTimeEntries should not return an error if there are no time entries but returned this: %s", err)
		} else {
			t.Logf("GetTimeEntries should return an empty list if the API returns no time entries")
		}

	}
}

func Test_GetTimeEntries_HTTPMethodIsGET(t *testing.T) {
	// arrange
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			// assert
			if method != "GET" {
				t.Fail()
				t.Logf("GetTimeEntries should use a GET request but didn't")
			}

			return []byte(``), nil
		},
	}

	timeEntryAPI := &TimeEntryAPI{
		restClient:    restClient,
		dateFormatter: date.NewISO8601Formatter(),
	}

	start := time.Date(2016, 1, 1, 0, 0, 1, 0, time.UTC)
	end := time.Date(2016, 6, 30, 23, 59, 59, 0, time.UTC)

	// act
	timeEntryAPI.GetTimeEntries(start, end)
}

func Test_GetTimeEntries_ValidJSONIsReturned_ProjectsAreReturned(t *testing.T) {
	// arrange
	timeEntriesJSON := `[
	{
		"id": 1,
		"wid": 1,
		"pid": 1,
		"billable": false,
		"start": "2016-09-06T06:33:56+00:00",
		"stop": "2016-09-06T06:48:51+00:00",
		"duration": 895,
		"description": "Lorem Ipsum"
	},
	{
		"id": 2,
		"guid": "fa4b0a79-8f76-437e-9675-de69c61206d9",
		"wid": 1,
		"pid": 2,
		"billable": false,
		"start": "2016-09-06T06:48:51+00:00",
		"stop": "2016-09-06T07:31:22+00:00",
		"description": "Yada Yada"
	}
]`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(timeEntriesJSON), nil
		},
	}

	timeEntryAPI := &TimeEntryAPI{
		restClient:    restClient,
		dateFormatter: date.NewISO8601Formatter(),
	}

	start := time.Date(2016, 1, 1, 0, 0, 1, 0, time.UTC)
	end := time.Date(2016, 6, 30, 23, 59, 59, 0, time.UTC)

	// act
	timeEntries, err := timeEntryAPI.GetTimeEntries(start, end)

	// assert
	if timeEntries == nil || len(timeEntries) != 2 {
		t.Fail()

		if err != nil {
			t.Logf("GetTimeEntries should have returned 2 time entries but returned an error instead: %s", err)
		} else {
			t.Logf("GetTimeEntries should have returned 2 time entries")
		}
	}
}
