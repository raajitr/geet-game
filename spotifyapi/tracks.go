package spotifyapi

import (
	"encoding/json"
	// "fmt"
	"net/url"
	"regexp"

	"github.com/tidwall/gjson"
)

type artists struct {
	Name string `json:"name"`
}

type externalUrl struct {
	SpotifyLink string `json:"spotify"`
}

type tracks struct {
	PreviewUrl string       `json:"preview_url"`
	Name       string       `json:"name"`
	AlbumName  string       `json:"album_name"`
	Artists    []artists    `json:"artists"`
	           externalUrl  `json:"external_urls"`
}

func (t *tracks) updateTrackAndAlbum() {
	var re = regexp.MustCompile(`(?m)(.*) \(From\s+\"(.*?)\"\)`)
	var str = &t.Name

	matches := re.FindAllStringSubmatch(*str, -1)

	for _, match := range matches {
		t.Name = match[1]
		t.AlbumName = match[2]
	}
}

func FetchTracks() []*tracks {
	api := urlBuilder("https://api.spotify.com", "/v1/search")

	var queryParams = map[string]string {
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

	data := createRequestData(queryParams)
	api.RawQuery = data.Encode()

	req := newRequest("GET", api.String(), url.Values{}, headers)

	respBody := makeRequest(req)

	result := gjson.GetBytes(respBody, "tracks.items")
	items := json.RawMessage(result.String())
	
	var trackList []*tracks
	json.Unmarshal(items, &trackList)

	for _, i := range trackList {
		i.updateTrackAndAlbum()
	}

	return trackList
}
