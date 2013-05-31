package github

import (
	"fmt"
	"net/http"
)

func FindUser(username string) (User, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s", username))

	if err != nil {
		return User{}, err
	}

	data, err := parseBody(resp.Body)
	resp.Body.Close()

	if err != nil {
		return User{}, err
	}

	j, err := parseJson(data)

	if err != nil {
		return User{}, err
	}

	return createUser(j), err
}
