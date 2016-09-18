package togglapi

import "testing"

func Test_NewTimeEntryRepository(t *testing.T) {
	// act
	repository := NewTimeEntryRepository("http://api.example.com", "sakldjaksljkl312312")

	// assert
	if repository == nil {
		t.Fail()
		t.Logf("NewTimeEntryRepository should have returned a time entry API client")
	}
}
