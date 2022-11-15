package main

import (
	"encoding/json"
	"fmt"
	"gotweet/pkg/sdk"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func init() {
	// setup authentication configuration
	viper.SetConfigName("auth")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		username := r.URL.Query().Get("username")
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
	log.Println("Server listening on port 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
