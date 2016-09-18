package togglapi

import "testing"

func Test_NewProjectRepository(t *testing.T) {
	// act
	repository := NewProjectRepository("http://api.example.com", "sakldjaksljkl312312")

	// assert
	if repository == nil {
		t.Fail()
		t.Logf("NewProjectRepository should have returned a project API client")
	}
}
