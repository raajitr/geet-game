package spotifyapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type tracks struct {
	Href string `json:"href"`
}

func FetchTracks() {
	api := urlBuilder("https://api.spotify.com", "/v1/search")

	var reqData = map[string]string {
		"q": "genre:bollywood",
		"type": "track",
		"market": "IN",
		"limit": "3",
		"offset": "2",
	}

	auth := spotifyAuth()
	var headers = map[string]string {
		"Authorization": "Bearer " + auth.AccessToken,
		"Content-Type":  "application/json",
	}

	data := createRequestData(reqData)
	api.RawQuery = data.Encode()

	req := newRequest("GET", api.String(), url.Values{}, headers)

	respBody := makeRequest(req)
	// fmt.Println(string(respBody))
	
	var resp tracks
	json.Unmarshal(respBody, &resp)

	fmt.Println(resp.Href)
	fmt.Println(resp)

	// return resp
}
