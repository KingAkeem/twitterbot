package main

import (
	"bytes"
	"fmt"
	"gotweet/pkg/api"
	"gotweet/pkg/auth"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("auth.json")
	viper.Set("ApiToken", "AAAAAAAAAAAAAAAAAAAAAOoQjAEAAAAAL035N2YaEV6NUTxvT%2BJbmnxbM9s%3DpBeFuGYvsnC2FoqY4nc6KPFTbvHbTLv0Awwkh449ga6afpDR3S")
}

func main() {
	clientID := "Ql9hX0ZBV1FWYnRxTXZLbDFqdDU6MTpjaQ"
	auth.AuthorizeUser(clientID, "http://localhost:8082")

	me, err := api.LookupUser("MynameisAkeem")
	if err != nil {
		panic(err)
	}
	_, err = api.ListFollowers(me.ID)
	if err != nil {
		panic(err)
	}

	user, err := api.LookupUser("Letrell57093938")
	if err != nil {
		panic(err)
	}

	// create JSON payload
	var jsonStr = []byte(fmt.Sprintf(`{"message": {"text": "Hello to just you two, this is a new group conversation."},"participant_ids": ["%s", "%s"],"conversation_type": "Group"}`, me.ID, user.ID))
	fmt.Println(string(jsonStr))
	req, err := http.NewRequest("POST", "https://api.twitter.com/2/dm_conversations", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	// set access token
	req.Header.Add("Authorization", "Bearer "+viper.GetString("AccessToken"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("response Body:", string(body))

}
