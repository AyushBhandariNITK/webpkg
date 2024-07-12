package restclient

import "context"

type NullRequest struct {
}

var nullRequest *NullRequest

func NewNullRequest() *NullRequest {
	return &NullRequest{}
}

func (r *NullRequest) SendRequest(ctx context.Context) error {
	return nil
}
