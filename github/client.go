package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	client *http.Client
	domain string
	Users  *UsersService
}

func NewClient() *Client {
	c := &Client{
		client: http.DefaultClient,
		domain: "https://api.github.com",
	}

	c.Users = &UsersService{
		client:   c,
		basePath: "users",
	}

	return c
}

func (c *Client) Get(path string, v interface{}) (*http.Response, error) {
	resp, err := c.client.Get(c.formatUrl(path))
	defer resp.Body.Close()

	if err != nil {
		return resp, err
	}

	err = json.NewDecoder(resp.Body).Decode(v)

	return resp, err
}

func (c *Client) formatUrl(path string) string {
	return fmt.Sprintf("%s/%s", c.domain, path)
}
