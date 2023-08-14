package spotifyapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

type authResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func generateB64Credentials() string {
	client_id := os.Getenv("SPOTIFY_CLIENT_ID")
	client_secret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	auth_vals := client_id + ":" + client_secret
    fmt.Println(auth_vals)

	return base64.StdEncoding.EncodeToString([]byte(auth_vals))
}

func spotifyAuth() *authResp {	
	api := urlBuilder("https://accounts.spotify.com", "/api/token")
	
	var reqData = map[string]string {
		"grant_type": "client_credentials",
	}

	clientCredenitals := generateB64Credentials()
	var headers = map[string]string {
		"Authorization": "Basic " + clientCredenitals,
		"Content-Type":  "application/x-www-form-urlencoded",
	}

	data := createRequestData(reqData)
	req := newRequest("POST", api.String(), data, headers)

	respBody := makeRequest(req)
	
	var resp *authResp
	json.Unmarshal(respBody, &resp)

	fmt.Println(resp.TokenType)

	return resp
}
