package api

import (
	"net/http"
	"time"
)

type Client struct {
	cl     *http.Client
	token  string
	apiURL string
}

func NewClient(token, apiURL string, timeout time.Duration) *Client {
	return &Client{token: token, cl: &http.Client{Timeout: timeout}, apiURL: apiURL}
}
