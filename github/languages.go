package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LanguageDetails struct {
	Count   int
	Percent float64
}

type Languages map[string]*LanguageDetails

type LanguagesService struct {
	client *Client
}

func (s *LanguagesService) Get(url string) (languages *Languages, err error) {
	languages = new(Languages)
	var resp *http.Response
	var data []byte
	var languageJson interface{}

	resp, err = s.client.GetUrlRaw(url)
	defer resp.Body.Close()

	if err != nil {
		return
	}

	data, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &languageJson)

	if err != nil {
		return
	}

	for l, c := range languageJson.(map[string]interface{}) {
		fmt.Println(l)
		fmt.Println(int(c.(float64)))
		//languages[l] = LanguageDetails{Count: c.(int)}
	}

	return
}
