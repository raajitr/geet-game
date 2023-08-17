package main

import (
	// "encoding/json"
	"fmt"
	spotifyapi "geet-game/spotifyapi"
	"net/http"

	ws "geet-game/ws"
	// "log"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func trackList(c *gin.Context) {
	trackList := spotifyapi.FetchTracks()
	c.JSON(http.StatusOK, trackList)
}

func setupRoute() {
	manager := ws.NewManager()
	router := gin.Default()

	router.GET("/", homePage)
	router.GET("/tracklist", trackList)
	router.GET("/ws/:id", manager.ServeWS)
	// http.HandleFunc("/ws/room", manager.ServeWS)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	router.Run(":8080")
}

func main() {
	fmt.Println("Websocket test")
	setupRoute()
}
