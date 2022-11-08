package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotweet/pkg/conf"
	"gotweet/pkg/store"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func SendGroupDm(text string, ids []string) {
	// create JSON payload
	var jsonStr = []byte(fmt.Sprintf(`{"message": {"text": "%s"},"participant_ids": ["%+v"],"conversation_type": "Group"}`, text, strings.Join(ids, ",")))
	fmt.Println(string(jsonStr))
	req, err := http.NewRequest("POST", "https://api.twitter.com/2/dm_conversations", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	// set access token
	req.Header.Add("Authorization", "Bearer "+viper.GetString("AccessToken"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	/*
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
	*/
}

// ListFollowers returns a list of users that follow the user with the given ID,
// an access token is required to perform this request
func ListFollowers(id string) ([]store.User, error) {
	url := fmt.Sprintf("%s/users/%s/followers", conf.BaseURL, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// inject access token using value saved in config file
	req.Header.Add("Authorization", "Bearer "+viper.GetString("AccessToken"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// parse and return response data
	result := struct {
		Data []store.User `json:"data"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// LookupUserByUsername returns the user with the given username
func LookupUserByUsername(username string) (*store.User, error) {
	url := fmt.Sprintf("%s/users/by/username/%s", conf.BaseURL, username)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// inject access or api token using value saved in config file
	req.Header.Add("Authorization", "Bearer "+viper.GetString("AccessToken"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// parse and return response data
	result := struct {
		Data store.User `json:"data"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}
