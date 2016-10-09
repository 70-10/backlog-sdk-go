package backlog

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultTimeout = 30 * time.Second
)

type Client struct {
	BaseURL *url.URL
	APIKey  string
	Timeout time.Duration
}

func NewClient(apiKey, spaceKey string) *Client {
	u, _ := url.Parse(fmt.Sprintf("https://%s.backlog.jp/api/v2", spaceKey))
	return &Client{
		BaseURL: u,
		APIKey:  apiKey,
		Timeout: defaultTimeout,
	}
}

func (c *Client) Get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.urlFor(path).String(), nil)

	if err != nil {
		return nil, err
	}

	return c.request(req)
}

func (c *Client) urlFor(path string) *url.URL {
	newURL, err := url.Parse(c.BaseURL.String())
	if err != nil {
		panic("invalid url passed")
	}
	newURL.Path += path
	return newURL
}

func (c *Client) request(req *http.Request) (*http.Response, error) {
	c.buildRequest(req)

	client := &http.Client{}
	client.Timeout = c.Timeout
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp, fmt.Errorf("API result failed: %s", resp.Status)
	}
	return resp, err
}

func (c *Client) buildRequest(req *http.Request) {
	q := req.URL.Query()
	q.Add("apiKey", c.APIKey)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Content-Type", "application/json")
}
