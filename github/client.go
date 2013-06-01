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
		repos:    &ReposService{client: c},
	}

	return c
}

func (c *Client) GetUrl(url string, v interface{}) (*http.Response, error) {
	resp, err := c.client.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return resp, err
	}

	err = json.NewDecoder(resp.Body).Decode(v)

	return resp, err
}

func (c *Client) GetPath(path string, v interface{}) (*http.Response, error) {
	return c.GetUrl(c.formatUrl(path), v)
}

func (c *Client) formatUrl(path string) string {
	return fmt.Sprintf("%s/%s", c.domain, path)
}
