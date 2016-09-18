package togglapi

import "testing"

func Test_NewAPI(t *testing.T) {
	// act
	client := NewAPI("http://api.example.com", "sakldjaksljkl312312")

	// assert
	if client == nil {
		t.Fail()
		t.Logf("NewAPI should have returned a Toggl API client")
	}
}
