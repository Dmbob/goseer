package main

import (
	"fmt"
	"net/http"
)

func main() {
	PORT := "8080" // Port for the webserver to listen, default 8080.

	//initDB()

	// Initialize the Websocket endpoint
	http.HandleFunc("/", connectWebsocket)
	fmt.Println("Server started and listening on port "+PORT)
	http.ListenAndServe(":"+PORT, nil)
}
