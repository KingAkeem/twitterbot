package sdk

import (
	"fmt"
	"gotweet/pkg/store"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLookupUserByUsername(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	username := "test"
	url := fmt.Sprintf("%s/users/by/username/%s?user.fields=%s", viper.GetString("BASE_URL"), username, strings.Join(user_fields, ","))

	createdAt := time.Now()
	expectedUser := store.User{
		ID:        "user-id",
		Username:  username,
		CreatedAt: createdAt,
		Verified:  false,
	}
	responder, err := httpmock.NewJsonResponder(http.StatusOK,
		struct {
			Data store.User
		}{
			Data: expectedUser,
		})
	assert.Nil(t, err)

	httpmock.RegisterResponder(http.MethodGet, url, responder)

	user, err := LookupUserByUsername(username)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, expectedUser.ID)
	assert.Equal(t, user.Username, expectedUser.Username)
	assert.False(t, user.Verified)
}

func TestListFollowers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "user-id"
	url := fmt.Sprintf("%s/users/%s/followers?user.fields=%s", viper.GetString("BASE_URL"), id, strings.Join(user_fields, ","))

	createdAt := time.Now()
	expectedUsers := []store.User{
		{
			ID:        id,
			Username:  "test user",
			CreatedAt: createdAt,
			Verified:  false,
		},
	}
	responder, err := httpmock.NewJsonResponder(http.StatusOK,
		struct {
			Data []store.User
		}{
			Data: expectedUsers,
		})
	assert.Nil(t, err)

	httpmock.RegisterResponder(http.MethodGet, url, responder)

	users, err := ListFollowers(id)
	assert.Nil(t, err)
	assert.Len(t, users, len(expectedUsers))
	assert.Equal(t, users[0].ID, expectedUsers[0].ID)
	assert.Equal(t, users[0].Username, expectedUsers[0].Username)
	assert.False(t, users[0].Verified)
}

func TestListFollowing(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "user-id"
	url := fmt.Sprintf("%s/users/%s/following?user.fields=%s", viper.GetString("BASE_URL"), id, strings.Join(user_fields, ","))

	createdAt := time.Now()
	expectedUsers := []store.User{
		{
			ID:        id,
			Username:  "test user",
			CreatedAt: createdAt,
			Verified:  false,
		},
	}
	responder, err := httpmock.NewJsonResponder(http.StatusOK,
		struct {
			Data []store.User
		}{
			Data: expectedUsers,
		})
	assert.Nil(t, err)

	httpmock.RegisterResponder(http.MethodGet, url, responder)

	users, err := ListFollowing(id)
	assert.Nil(t, err)
	assert.Len(t, users, len(expectedUsers))
	assert.Equal(t, users[0].ID, expectedUsers[0].ID)
	assert.Equal(t, users[0].Username, expectedUsers[0].Username)
	assert.False(t, users[0].Verified)
}

func TestListTweets(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	username := "test"
	testAPI := viper.GetString("BASE_URL")
	url := fmt.Sprintf(`%s/tweets/search/recent?query=from:%s&tweet.fields=%s`, testAPI, username, strings.Join(tweet_fields, ","))

	createdAt := time.Now()
	expectedTweets := []store.Tweet{
		{
			ID:        "test-id",
			Text:      "this is a test tweet",
			CreatedAt: createdAt,
		},
	}
	responder, err := httpmock.NewJsonResponder(http.StatusOK,
		struct {
			Data []store.Tweet
		}{
			Data: expectedTweets,
		})
	assert.Nil(t, err)

	httpmock.RegisterResponder(http.MethodGet, url, responder)

	tweets, err := ListTweets("test")
	assert.Nil(t, err)
	assert.Len(t, tweets, len(expectedTweets))
	assert.Equal(t, tweets[0].ID, expectedTweets[0].ID)
	assert.Equal(t, tweets[0].Text, expectedTweets[0].Text)
}
