package restclient

import (
	"context"
	"webpkg/log"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	BaseURL  string
	Port     int
	Endpoint string
	client   *resty.Client
}

func (c *Client) Builder() *Client {
	return &Client{}
}

func NewGETClient() *Client {
	return &Client{
		client: resty.New(),
	}
}

func (c *Client) SetBaseURL(url string) *Client {
	c.BaseURL = url
	return c
}

func (c *Client) SetEndpoint(endpoint string) *Client {
	c.Endpoint = endpoint
	return c
}

func (c *Client) SetPort(port int) *Client {
	c.Port = port
	return c
}

func (c *Client) SetHeaders(opts ...map[string]string) *Client {
	if len(opts) > 0 {
		c.client.SetHeaders(opts[0])
	} else {
		c.client.SetHeaders(defaultHeaders)
	}
	return c
}

func (c *Client) Get(enpoint string) (*resty.Response, error) {
	return c.client.R().Get(enpoint)
}

func (client *Client) SendRequest(ctx context.Context) error {
	resp, err := client.
		SetHeaders(defaultHeaders).
		Get(client.Endpoint)

	if err != nil {
		log.Print(log.Error, "Error: %s", err.Error())
		return err
	}
	log.Print(log.Info, "Request response : %+v", resp)
	return nil
}
