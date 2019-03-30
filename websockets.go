package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
)

// Global Variables are in all caps.
var CLIENTS = make(map[*websocket.Conn]bool)
var BROADCAST = make(chan Message)
var UPGRADER = websocket.Upgrader{}
var STATE = "PAUSED"
var TIME = "0"

// Define our message object
type Message struct {
        Event string `json:"event"`
        Value float32 `json:"value"`
        //Date string `json:"date"`
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	// convert to WS, store in client global.
        ws, err := UPGRADER.Upgrade(w, r, nil)
        if err != nil {
                fmt.Println(err)
        }
        defer ws.Close()
        //CLIENTS[ws] = true

	// Read message if there is one.
	for {
		var m Message
		err := ws.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}
		fmt.Println(m)
	}
}
