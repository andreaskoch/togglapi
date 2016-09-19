package togglapi

import "testing"

func Test_NewClientAPI(t *testing.T) {
	// act
	api := NewClientAPI("http://api.example.com", "sakldjaksljkl312312")

	// assert
	if api == nil {
		t.Fail()
		t.Logf("NewClientAPI should have returned a project API client")
	}
}
