package main

import (
	"encoding/json"
	"fmt"
	"gotweet/pkg/conf"
	"gotweet/pkg/sdk"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	// attempt to file env file and load config
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		// find twitterbot.env file and load config from directory
		if info.Name() == "twitterbot.env" {
			conf.LoadConfig(strings.Replace(path, "twitterbot.env", "", 1))
		}
		return nil
	})

	router := mux.NewRouter()
	router.HandleFunc("/tweets/{username}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		username := mux.Vars(r)["username"]
		tweets, err := sdk.ListTweets(username)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to lookup followers by ID %s. Error: %+v", username, err)
			log.Println(errMsg)
			w.Write([]byte(errMsg))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(tweets)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to write followers %+v. Error: %+v", tweets, err)
			log.Println(errMsg)
			w.Write([]byte("Unable to write user object"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

	router.HandleFunc("/following/{username}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		username := mux.Vars(r)["username"]
		user, err := sdk.LookupUserByUsername(username)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to lookup followers by ID %s. Error: %+v", user.ID, err)
			log.Println(errMsg)
			w.Write([]byte(errMsg))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		followers, err := sdk.ListFollowing(user.ID)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to lookup followers by ID %s. Error: %+v", user.ID, err)
			log.Println(errMsg)
			w.Write([]byte(errMsg))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(followers)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to write followers %+v. Error: %+v", followers, err)
			log.Println(errMsg)
			w.Write([]byte("Unable to write user object"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

	router.HandleFunc("/followers/{username}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		username := mux.Vars(r)["username"]
		user, err := sdk.LookupUserByUsername(username)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to lookup followers by ID %s. Error: %+v", user.ID, err)
			log.Println(errMsg)
			w.Write([]byte(errMsg))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		followers, err := sdk.ListFollowers(user.ID)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to lookup followers by ID %s. Error: %+v", user.ID, err)
			log.Println(errMsg)
			w.Write([]byte(errMsg))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(followers)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to write followers %+v. Error: %+v", followers, err)
			log.Println(errMsg)
			w.Write([]byte("Unable to write user object"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

	router.HandleFunc("/user/{username}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		username := mux.Vars(r)["username"]
		user, err := sdk.LookupUserByUsername(username)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to lookup user by username %s. Error: %+v", username, err)
			log.Println(errMsg)
			w.Write([]byte(errMsg))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to write user %+v. Error: %+v", user, err)
			log.Println(errMsg)
			w.Write([]byte("Unable to write user object"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

	port := viper.GetString("PORT")
	log.Println("Server listening on port " + port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		panic(err)
	}
}
