package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/websocket"
)

type Message struct {
	IsRequest bool
	RequestType string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func parseMessage(msg string) Message {
	var message Message
	json.Unmarshal([]byte(msg), &message)

	return message
}

func wsReader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("Message Type: "+string(messageType))
			fmt.Println(err)
			return
		}

		var messageRecieved Message
		messageRecieved = parseMessage(string(p))

		if messageRecieved.IsRequest {
			if messageRecieved.RequestType == "systemtemps" {
				temps := fetchSystemTemps()
				conn.WriteJSON(temps)
			}
		}
	}
}

func connectWebsocket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Client connected successfully...")

	wsReader(ws)
}
