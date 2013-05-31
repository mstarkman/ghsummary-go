package github

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func parseJson(j []byte) (map[string]interface{}, error) {
	var result interface{}

	err := json.Unmarshal(j, &result)

	return result.(map[string]interface{}), err
}

func createUser(j map[string]interface{}) User {
	return User{
		Name: j["name"].(string),
	}
}

func parseBody(body io.ReadCloser) ([]byte, error) {
	return ioutil.ReadAll(body)
}

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
