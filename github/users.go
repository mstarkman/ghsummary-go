package github

import (
	"fmt"
)

type User struct {
	Name            string `json:"name,omitempty"`
	Login           string `json:"login,omitempty"`
	PublicRepoCount int    `json:"public_repos,omitempty"`
}

type UsersService struct {
	client   *Client
	basePath string
}

func (s *UsersService) Get(username *string) (*User, error) {
	path := fmt.Sprintf("%s/%s", s.basePath, *username)
	user := new(User)

	_, err := s.client.Get(path, user)

	return user, err
}
