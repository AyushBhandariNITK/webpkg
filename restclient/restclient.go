package restclient

import (
	"context"
	"errors"
	"fmt"
	"webpkg/log"
)

type Request interface {
	SendRequest(context.Context) error
}

type RequestFactory struct {
	Request
}

var RequestMap = map[string]Request{
	"GET": NewGETClient(),
}

func NewRequest(RequestType string) (Request, error) {
	log.Print(log.Info, "Request Type %s", RequestType)
	if val, ok := RequestMap[RequestType]; ok {
		return val, nil
	} else {
		errMsg := fmt.Sprintf("Request Type %s is not supported", RequestType)
		return nullRequest, errors.New(errMsg)
	}

}

var defaultHeaders map[string]string

func init() {
	defaultHeaders = map[string]string{
		"Content-type": "Application/json",
		"Accept":       "Application/json",
	}
	nullRequest = NewNullRequest()

}
