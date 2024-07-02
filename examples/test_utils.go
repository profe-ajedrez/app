package examples

import (
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	do         func()
}

func NewClient(apiURL string) *Client {
	return &Client{
		baseURL: apiURL,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *Client) WithDo(d func()) {
	c.do = d
}
