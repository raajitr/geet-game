package spotifyapi

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func urlBuilder(host string, resource string) *url.URL {
    u, _ := url.ParseRequestURI(host)
    u.Path = resource

	return u
}

func createRequestData(reqData map[string]string) url.Values {
	data := url.Values{}
	for key, val := range reqData {
		data.Add(key, val)
	}

	return data
}

func newRequest(method string, api string, data url.Values, headers map[string]string) *http.Request {
	r, _ := http.NewRequest(method, api, strings.NewReader(data.Encode())) // URL-encoded payload

	for key, val := range headers {
		r.Header.Add(key, val)
	}

	return r
}

func makeRequest(req *http.Request) []byte {
	client := &http.Client{}
    resp, _ := client.Do(req)

    body, _ := io.ReadAll(resp.Body)

	return body
}
