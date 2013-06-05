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

	c.Users.repos.languages = &LanguagesService{client: c}

	return c
}

func (c *Client) GetUrlRaw(url string) (resp *http.Response, err error) {
	resp, err = c.client.Get(url)
	return

}

func (c *Client) GetUrl(url string, v interface{}) (resp *http.Response, err error) {
	resp, err = c.GetUrlRaw(url)
	defer resp.Body.Close()

	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	return
}

func (c *Client) GetPath(path string, v interface{}) (*http.Response, error) {
	return c.GetUrl(c.formatUrl(path), v)
}

func (c *Client) formatUrl(path string) string {
	return fmt.Sprintf("%s/%s", c.domain, path)
}
