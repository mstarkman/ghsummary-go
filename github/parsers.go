package github

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func parseBody(body io.ReadCloser) ([]byte, error) {
	return ioutil.ReadAll(body)
}

func parseJson(j []byte) (map[string]interface{}, error) {
	var result interface{}

	err := json.Unmarshal(j, &result)

	return result.(map[string]interface{}), err
}
