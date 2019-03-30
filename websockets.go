package main

import (
	"net/http"
	"fmt"
	//"strings"
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
        Event string `json:"Event"`
        Value float32 `json:"Value"`
	URL string `json:"URL"`
        Date int `json:"Date"`
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	// convert to WS, store in client global.
        ws, err := UPGRADER.Upgrade(w, r, nil)
        if err != nil {
                fmt.Println(err)
        }
        //defer ws.Close()
        CLIENTS[ws] = true
	// Read message if there is one.
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}
		fmt.Println("Message received! :: " + msg.URL)
		/*if(strings.Compare(msg.Event, "PAUSED") == 0){
			broadcast(msg);
		}else if(strings.Compare(msg.Event, "PLAYING") == 0){
			broadcast(msg);
		}else if(strings.Compare(msg.Event, "BUFFER") == 0){
			broadcast(msg);
		}else if(strings.Compare(msg.Event, "VIDEO") == 0){
			broadcast(msg);
		}*/
		// Just gonna remove because without validation code is much simpler, thus better.
		broadcast(msg, ws);
	}
	//ws.WriteJSON("{ video:" + VIDEO + ", time:" + TIME + ", state:" + STATE +" }")
}

func broadcast(msg Message, sender *websocket.Conn){
       for client := range CLIENTS {
              if (client != sender){
                    err := client.WriteJSON(msg)
                    if err != nil {
                           fmt.Printf("error: %v", err)
                           client.Close()
                           delete(CLIENTS, client)
                    }
             }
       }
}