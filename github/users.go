package github

import (
	"fmt"
	"time"
)

type User struct {
	LoadTime        float64
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

func (s *UsersService) Get(username *string) (user *User, err error) {
	startTime := time.Now()
	path := fmt.Sprintf("%s/%s", s.basePath, *username)
	user = new(User)

	_, err = s.client.GetPath(path, user)
	user.Repos, err = s.getReposDetails(user)

	user.LoadTime = time.Since(startTime).Seconds()

	return
}

func (s *UsersService) getReposDetails(u *User) (repos *[]Repo, err error) {
	repos, err = s.repos.List(u.ReposUrl)
	return
}
