package togglapi

import (
	"fmt"
	"io"
	"testing"

	"github.com/andreaskoch/togglapi/date"
	"github.com/andreaskoch/togglapi/model"
)

func Test_CreateTimeEntry_RestClientReturnsError_ErrorIsReturned(t *testing.T) {
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

	input := model.TimeEntry{}

	// act
	_, err := timeEntryAPI.CreateTimeEntry(input)

	// assert
	if err == nil {
		t.Fail()
		t.Logf("CreateTimeEntry should return an error if the REST client returned an error")
	}
}

func Test_CreateTimeEntry_InvalidJSONIsReturned_ErrorIsReturned(t *testing.T) {
	// arrange
	timeEntryJSON := `[
  {;,,,.,daskdlasdlak ---invalid--
    "id": 1,
]`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(timeEntryJSON), nil
		},
	}

	timeEntryAPI := &TimeEntryAPI{
		restClient:    restClient,
		dateFormatter: date.NewISO8601Formatter(),
	}

	input := model.TimeEntry{}

	// act
	_, err := timeEntryAPI.CreateTimeEntry(input)

	// assert
	if err == nil {
		t.Fail()
		t.Logf("CreateTimeEntry should return an error if the JSON returned by the API is invalid")
	}
}

func Test_CreateTimeEntry_NoTimeEntryReturned_EmptyTimeEntryIsReturned(t *testing.T) {
	// arrange
	timeEntryJSON := `null`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(timeEntryJSON), nil
		},
	}

	timeEntryAPI := &TimeEntryAPI{
		restClient:    restClient,
		dateFormatter: date.NewISO8601Formatter(),
	}

	input := model.TimeEntry{}

	// act
	_, err := timeEntryAPI.CreateTimeEntry(input)

	// assert
	if err != nil {
		t.Fail()
		t.Logf("CreateTimeEntry should return an empty time entry if the API returns no time entry but returned an error instead: %s", err)

	}
}

func Test_CreateTimeEntry_HTTPMethodIsPOST(t *testing.T) {
	// arrange
	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			// assert
			if method != "POST" {
				t.Fail()
				t.Logf("CreateTimeEntry should use POST method but didn't")
			}

			return nil, nil
		},
	}

	timeEntryAPI := &TimeEntryAPI{
		restClient:    restClient,
		dateFormatter: date.NewISO8601Formatter(),
	}

	input := model.TimeEntry{}

	// act
	timeEntryAPI.CreateTimeEntry(input)
}

func Test_CreateTimeEntry_ValidJSONIsReturned_ProjectsAreReturned(t *testing.T) {
	// arrange
	timeEntryJSON := `{
	"data": {
		"id": 1,
		"wid": 1,
		"pid": 1,
		"billable": false,
		"start": "2016-09-06T06:33:56+00:00",
		"stop": "2016-09-06T06:48:51+00:00",
		"description": "Lorem Ipsum"
	}
}`

	restClient := &mockRESTRequester{
		request: func(method, route string, payload io.Reader) ([]byte, error) {
			return []byte(timeEntryJSON), nil
		},
	}

	timeEntryAPI := &TimeEntryAPI{
		restClient:    restClient,
		dateFormatter: date.NewISO8601Formatter(),
	}

	input := model.TimeEntry{}

	// act
	timeEntry, err := timeEntryAPI.CreateTimeEntry(input)

	// assert
	if err != nil || timeEntry.Description != "Lorem Ipsum" {
		t.Fail()

		if err != nil {
			t.Logf("CreateTimeEntry should have returned 2 time entry but returned an error instead: %s", err)
		} else {
			t.Logf("CreateTimeEntry should have returned 2 time entry")
		}
	}
}
