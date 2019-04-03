package main

import (
	"net/http"
	"fmt"
	"strings"
	"github.com/gorilla/websocket"
)

// Global Variables are in all caps.
var CLIENTS = make(map[*websocket.Conn]bool)
var BROADCAST = make(chan Message)
var UPGRADER = websocket.Upgrader{}
var STATE = "PAUSED"
var TIME = float32(0)
var VIDEO = "http://www.youtube.com/v/dQw4w9WgXcQ?version=3Message"

// Define our message object
type Message struct {
        Event string `json:"Event"`
        Value float32 `json:"Value"`
	URL string `json:"URL"`
        Date int `json:"Date"`
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	// convert to WS, store in client global.
        ws, err := UPGRADER.Upgrade(w, r, w.Header())
        if err != nil {
                fmt.Println(err)
        }

	// Gonna check if they are in the list, if not, say hello and register them.
	if (!CLIENTS[ws]){
		defer ws.Close()
		CLIENTS[ws] = true
		fmt.Println("Adding new connection!");
		// Need to know if he is not first, so we can sync everyone.
		if(len(CLIENTS) > 0){
			fmt.Println("Syncing!");
			broadcast(Message{"SYNC",float32(TIME),VIDEO,0}, ws)
		}
		ws.WriteJSON(Message{"HELLO", float32(TIME), VIDEO, 0})
	}

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}
		//fmt.Println("Message received! :: " + msg.Event)
		if(strings.Compare(msg.Event, "PLAYING") == 0 || strings.Compare(msg.Event, "PAUSED") == 0){
			STATE = msg.Event
			broadcast(msg, ws)
		}else if(strings.Compare(msg.Event, "VIDEO") == 0){
			fmt.Println("Video Changing to :: " + msg.URL)
			VIDEO = msg.URL
			TIME = 0
			broadcast(msg, ws)
		}else if(strings.Compare(msg.Event, "UPDATE") == 0){
			//fmt.Println("Time Changing to :: " + string(msg.Value))
			TIME = msg.Value
			broadcast(msg, ws)
		}/*else if(strings.Compare(msg.Event, "UPDATE") == 0){
			ws.WriteJSON(Message{"UPDATE",float32(TIME),VIDEO,0})
		}*/
	}
}

func broadcast(msg Message, sender *websocket.Conn){
       for client := range CLIENTS {
              if(client != sender){
                    err := client.WriteJSON(msg)
                    if(err != nil){
                           fmt.Printf("error: %v", err)
                           client.Close()
                           delete(CLIENTS, client)
                    }
             }
       }
}

