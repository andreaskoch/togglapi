package togglapi

import "testing"

func Test_NewClientRepository(t *testing.T) {
	// act
	repository := NewClientRepository("http://api.example.com", "sakldjaksljkl312312")

	// assert
	if repository == nil {
		t.Fail()
		t.Logf("NewClientRepository should have returned a project API client")
	}
}
