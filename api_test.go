package togglapi

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_NewAPI(t *testing.T) {
	// act
	client := NewAPI("http://api.example.com", "sakldjaksljkl312312")

	// assert
	if client == nil {
		t.Fail()
		t.Logf("NewAPI should have returned a Toggl API client")
	}
}

// To initialize a new Toggl API instance use the NewAPI function and pass
// Toggl API URL (e.g. "https://www.toggl.com/api/v8") and your personal
// API token (e.g. "12jkhjh2j3jkj23").
func ExampleNewAPI() {
	baseURL := "https://www.toggl.com/api/v8"
	apiToken := "Your-API-Token"
	api := NewAPI(baseURL, apiToken)

	stop := time.Now()
	start := stop.AddDate(0, -1, 0)

	timeEntries, timeEntriesError := api.GetTimeEntries(start, stop)
	if timeEntriesError != nil {
		fmt.Fprintf(os.Stderr, "Failed to get timeEntries: %s", timeEntriesError)
		return
	}

	for _, timeEntry := range timeEntries {
		fmt.Printf("%s - %s: %s\n", timeEntry.Start, timeEntry.Stop, timeEntry.Description)
	}
}
