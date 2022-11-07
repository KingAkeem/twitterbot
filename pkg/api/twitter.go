package api

import (
	"encoding/json"
	"fmt"
	"gotweet/pkg/conf"
	"gotweet/pkg/store"
	"net/http"

	"github.com/spf13/viper"
)

var client = &http.Client{}

func ListFollowers(id string) ([]store.User, error) {
	url := fmt.Sprintf("%s/users/2244994945/followers", conf.BaseURL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+viper.GetString("AccessToken"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result := struct {
		Data []store.User `json:"data"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

func LookupUser(username string) (*store.User, error) {
	url := fmt.Sprintf("%s/users/by/username/%s", conf.BaseURL, username)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+viper.GetString("AccessToken"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	result := struct {
		Data store.User `json:"data"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}
