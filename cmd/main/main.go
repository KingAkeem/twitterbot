package main

import (
	"gotweet/pkg/api"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("auth.json")
	viper.Set("ApiToken", "AAAAAAAAAAAAAAAAAAAAAOoQjAEAAAAAL035N2YaEV6NUTxvT%2BJbmnxbM9s%3DpBeFuGYvsnC2FoqY4nc6KPFTbvHbTLv0Awwkh449ga6afpDR3S")
}

func main() {
	/*
		clientID := "Ql9hX0ZBV1FWYnRxTXZLbDFqdDU6MTpjaQ"
		auth.AuthorizeUser(clientID, "http://localhost:8082")
	*/

	me, err := api.LookupUserByUsername("MynameisAkeem")
	if err != nil {
		panic(err)
	}

	user, err := api.LookupUserByUsername("Letrell57093938")
	if err != nil {
		panic(err)
	}

	api.SendGroupDm("testing", []string{me.ID, user.ID})
}
