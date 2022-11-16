package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotweet/pkg/store"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

var (
	// additional fields used in requests for each object type
	user_fields  = []string{"created_at", "description", "location", "profile_image_url", "url", "verified"}
	tweet_fields = []string{"created_at", "source", "text", "geo"}
)

// SendGroupDm sends a single message to multiple users using the IDs given.
func SendGroupDm(text string, ids []string) error {
	// create JSON payload
	var jsonStr = []byte(fmt.Sprintf(`{"message": {"text": "%s"},"participant_ids": ["%+v"],"conversation_type": "Group"}`, text, strings.Join(ids, ",")))
	req, err := http.NewRequest("POST", "https://api.twitter.com/2/dm_conversations", bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// set access token, required for non-public operations
	req.Header.Add("Authorization", "Bearer "+viper.GetString("AccessToken"))

	_, err = http.DefaultClient.Do(req)
	return err
}

// ListTweets returns 100 of the most recent tweets for the given user in the past seven days
func ListTweets(username string) ([]store.Tweet, error) {
	url := fmt.Sprintf(`%s/tweets/search/recent?query=from:%s&tweet.fields=%s`, viper.GetString("BASE_URL"), username, strings.Join(tweet_fields, ","))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// inject access token using value saved in config file
	req.Header.Add("Authorization", "Bearer "+viper.GetString("API_TOKEN"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// parse and return response data
	result := struct {
		Data []store.Tweet `json:"data"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// ListFollowing returns a list of following for the user with the given ID.
func ListFollowing(id string) ([]store.User, error) {
	url := fmt.Sprintf("%s/users/%s/following?user.fields=%s", viper.GetString("BASE_URL"), id, strings.Join(user_fields, ","))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// inject access token using value saved in config file
	req.Header.Add("Authorization", "Bearer "+viper.GetString("API_TOKEN"))
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

// ListFollowers returns a list of followers for the user with the given ID.
func ListFollowers(id string) ([]store.User, error) {
	url := fmt.Sprintf("%s/users/%s/followers?user.fields=%s", viper.GetString("BASE_URL"), id, strings.Join(user_fields, ","))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// inject access token using value saved in config file
	req.Header.Add("Authorization", "Bearer "+viper.GetString("API_TOKEN"))
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
	// additional information added to the results
	url := fmt.Sprintf("%s/users/by/username/%s?user.fields=%s", viper.GetString("BASE_URL"), username, strings.Join(user_fields, ","))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// inject access or api token using value saved in config file
	req.Header.Add("Authorization", "Bearer "+viper.GetString("API_TOKEN"))
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
