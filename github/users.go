package github

import (
	"fmt"
)

type User struct {
	Name            string `json:"name,omitempty"`
	Login           string `json:"login,omitempty"`
	PublicRepoCount int    `json:"public_repos,omitempty"`
	Repos           *[]Repo
	ReposUrl        string `json:"repos_url,omitempty"`
}

type UsersService struct {
	client   *Client
	basePath string
	repos    *ReposService
}

func (s *UsersService) Get(username *string) (*User, error) {
	path := fmt.Sprintf("%s/%s", s.basePath, *username)
	user := new(User)

	_, err := s.client.GetPath(path, user)
	user.Repos, err = s.repos.List(user.ReposUrl)

	return user, err
}
