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
var VIDEO = ""

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
        CLIENTS[ws] = true
	// Read message if there is one.
	for {
		var m Message
		err := ws.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}
		fmt.Println(m)
		// Here we need to decide what if any action to take based on event.
		//	Most likely broadcasting out to clients changes.
	}
	ws.WriteJSON("{ video:" + VIDEO + ", time:" + TIME + ", state:" + STATE +" }")
}
/*
func broadcast(msg string){
       for client := range CLIENTS {
                err := client.WriteJSON(msg)
                if err != nil {
                       fmt.Printf("error: %v", err)
                       client.Close()
                       delete(CLIENTS, client)
                }
       }
}*/