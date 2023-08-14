package main

import (
	"fmt"
	"log"
	"net/http"
	spotifyapi "websocket-server/spotifyapi"
	"websocket-server/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func room(w http.ResponseWriter, r *http.Request) {
	ws, _ := websocket.Upgrade(w, r)
	go websocket.Writer(ws)
}

func setupRoute() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/room", room)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Websocket test")
	// setupRoute()
	spotifyapi.FetchTracks()
}
