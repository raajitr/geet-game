package main

import (
	"encoding/json"
	"fmt"
	spotifyapi "geet-game/spotifyapi"
	"geet-game/websocket"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func trackList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	trackList := spotifyapi.FetchTracks()
	
	resp, _ := json.Marshal(trackList)
	fmt.Fprintf(w, string(resp))
}

func room(w http.ResponseWriter, r *http.Request) {
	ws, _ := websocket.Upgrade(w, r)
	go websocket.Writer(ws)
}

func setupRoute() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/tracklist", trackList)
	http.HandleFunc("/room", room)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Websocket test")
	setupRoute()
}
