package togglapi

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_NewTimeEntryAPI(t *testing.T) {
	// act
	repository := NewTimeEntryAPI("http://api.example.com", "sakldjaksljkl312312")

	// assert
	if repository == nil {
		t.Fail()
		t.Logf("NewTimeEntryAPI should have returned a time entry API client")
	}
}

// If you are only interested in the Time Entry API you can instantiate a
// TimeEntryAPI using the NewTimeEntryAPI function.
func ExampleNewTimeEntryAPI() {
	apiToken := "Your-Toggl-API-Token"
	baseURL := "https://www.toggl.com/api/v8"
	timeEntryRepository := NewTimeEntryAPI(baseURL, apiToken)

	stop := time.Now()
	start := stop.AddDate(0, -1, 0)

	timeEntries, timeEntriesError := timeEntryRepository.GetTimeEntries(start, stop)
	if timeEntriesError != nil {
		fmt.Fprintf(os.Stderr, "Failed to get timeEntries: %s", timeEntriesError)
		return
	}

	for _, timeEntry := range timeEntries {
		fmt.Printf("%s - %s: %s\n", timeEntry.Start, timeEntry.Stop, timeEntry.Description)
	}
}
