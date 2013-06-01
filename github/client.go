package github

import (
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

func (c *Client) Get(path string) (*http.Response, error) {
	return c.client.Get(c.formatUrl(path))
}

func (c *Client) formatUrl(path string) string {
	return fmt.Sprintf("%s/%s", c.domain, path)
}
