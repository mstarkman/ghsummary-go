package github

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name,omitempty"`
}

type UsersService struct {
	client   *Client
	basePath string
}

func (s *UsersService) Get(username *string) (*User, error) {
	path := fmt.Sprintf("%s/%s", s.basePath, *username)
	resp, err := s.client.Get(path)
	defer resp.Body.Close()

	if err != nil {
		return new(User), err
	}

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	return user, err
}
