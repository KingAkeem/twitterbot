package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/viper"
)

/*
	// Example of retrieving access token using oauth2
	clientID := "Ql9hX0ZBV1FWYnRxTXZLbDFqdDU6MTpjaQ"
	auth.AuthorizeUser(clientID, "http://localhost:8082")
*/
// AuthorizeUser implements the PKCE OAuth2 flow.
func AuthorizeUser(clientID string, redirectURL string) {
	codeChallenge := "challenge"
	scopes := "users.read%20dm.write%20follows.read%20dm.read%20tweet.read%20tweet.write%20offline.access"
	// construct the authorization URL (with Auth0 as the authorization provider)
	authorizationURL := fmt.Sprintf(
		"https://twitter.com/i/oauth2/authorize?"+
			"scope=%s"+
			"&state=state"+
			"&response_type=code&client_id=%s"+
			"&code_challenge=%s"+
			"&code_challenge_method=plain&redirect_uri=%s",
		scopes, clientID, codeChallenge, redirectURL)

	// start a web server to listen on a callback URL
	server := &http.Server{Addr: redirectURL}

	// define a handler that will get the authorization code, call the token endpoint, and close the HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// get the authorization code
		code := r.URL.Query().Get("code")
		if code == "" {
			fmt.Println("Url Param 'code' is missing")
			io.WriteString(w, "Error: could not find 'code' URL parameter\n")

			// close the HTTP server and return
			cleanup(server)
			return
		}

		// trade the authorization code and the code verifier for an access token
		token, err := getAccessToken(clientID, codeChallenge, code, redirectURL)
		if err != nil {
			fmt.Println("could not get access token")
			io.WriteString(w, "Error: could not retrieve access token\n")

			// close the HTTP server and return
			cleanup(server)
			return
		}

		// AccessToken can be used for private information
		viper.Set("AccessToken", token)
		err = viper.WriteConfigAs("auth.json")
		if err != nil {
			fmt.Println("could not write config file", err)
			io.WriteString(w, "Error: could not store access token\n")

			// close the HTTP server and return
			cleanup(server)
			return
		}

		// return an indication of success to the caller
		io.WriteString(w, `
		<html>
			<body>
				<h1>Login successful!</h1>
				<h2>You can close this window and return to the gotweet CLI.</h2>
			</body>
		</html>`)

		fmt.Println("Successfully logged into twitter API.")

		// close the HTTP server
		cleanup(server)
	})

	// parse the redirect URL for the port number
	u, err := url.Parse(redirectURL)
	if err != nil {
		fmt.Printf("bad redirect URL: %s\n", err)
		os.Exit(1)
	}

	// set up a listener on the redirect port
	port := fmt.Sprintf(":%s", u.Port())
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("can't listen to port %s: %s\n", port, err)
		os.Exit(1)
	}

	// open a browser window to the authorizationURL
	err = open.Start(authorizationURL)
	if err != nil {
		fmt.Printf(": can't open browser to URL %s: %s\n", authorizationURL, err)
		os.Exit(1)
	}

	// start the blocking web server loop
	// this will exit when the handler gets fired and calls server.Close()
	server.Serve(l)
}

// getAccessToken trades the authorization code retrieved from the first OAuth2 leg for an access token
func getAccessToken(clientID string, codeVerifier string, authorizationCode string, callbackURL string) (string, error) {
	// set the url and form-encoded data for the POST to the access token endpoint
	data := fmt.Sprintf(
		"grant_type=authorization_code&client_id=%s"+
			"&code_verifier=%s"+
			"&code=%s"+
			"&redirect_uri=%s",
		clientID, codeVerifier, authorizationCode, callbackURL)
	payload := strings.NewReader(data)

	// create the request and execute it
	req, _ := http.NewRequest("POST", viper.GetString("TOKEN_URL"), payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	// Encode username:password combination using base64 encoding as new bearer token to acquire access token
	clientSecret := "H-i2gZeR5z8ycheJPGNGTiL6JCtF0i-Krt0SR0Dq4S3wFTGGy1"
	newBearerToken := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
	req.Header.Add("Authorization", "Basic "+newBearerToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("HTTP error: %s", err)
		return "", err
	}

	// process the response
	defer res.Body.Close()
	var responseData map[string]interface{}
	body, _ := ioutil.ReadAll(res.Body)

	// unmarshal the json into a string map
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Printf("JSON error: %s", err)
		return "", err
	}

	// retrieve the access token out of the map, and return to caller
	accessToken := responseData["access_token"].(string)
	return accessToken, nil
}

// cleanup closes the HTTP server
func cleanup(server *http.Server) {
	// we run this as a goroutine so that this function falls through and
	// the socket to the browser gets flushed/closed before the server goes away
	go server.Close()
}
