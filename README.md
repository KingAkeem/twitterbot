# twitterbot

## About 
twitterbot is an OSINT program focused on the social media application Twitter.
Users can view a wide variety of information from various sites with an easy-to-use GUI, information includes:
- User profile information such as name, location, bio, etc.
- User site posts
- User's followers
- User's following

## Setup
This application requires a variety of software to run locally, this is the only option currently.
You'll need a Bearer API token, create a `twitterbot.env` file with `API_TOKEN` key as the bearer token.
Add the `BASE_URL` AND `TOKEN_URL` values from the example as well, place this file in the `conf` directory.

e.g. example `twitterbot.env`
```env
BASE_URL="https://api.twitter.com/2"
TOKEN_URL="https://api.twitter.com/2/oauth2/token"
API_TOKEN="ASsafdasdfJASI!23Sasdf"
```

Languages used are:
1. JavaScript, HTML & CSS (Front-end uses ReactJS)
2. Golang (Server)

### Front-end
1. Go to `front-end` directory
2. Run `npm start`

This will build the front-end code and create a local server at `localhost:3000`. The site should automatically load in your default browser.
If it does not, then go to `http://localhost:3000`. Hot reloading is activated so could changes are immediately applied.

### Back-end
1. Go to `cmd/main` directory
2. Run `go run main.go`

This will start the back-end at `localhost:8081`. This needs to be run in tandem with the front-end to provide the necessary REST API.
The port can be updated from `twitterbot.env` using the `PORT` key.
