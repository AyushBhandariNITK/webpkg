package restclient_test

import (
	"testing"
	"webpkg/restclient"
)

func TestGetRequestType(t *testing.T) {
	requestType := "GET"
	_, err := restclient.NewRequest(requestType)
	if err != nil {
		t.Errorf("expected PASS but failed due to %s", err.Error())
	}

}

func TestNonDefineRequestType(t *testing.T) {
	requestType := "Undefined"
	_, err := restclient.NewRequest(requestType)
	if err == nil {
		t.Errorf("expect FAIL but PASSED")
	}
}
