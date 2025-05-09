package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrade HTTP connection to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origins
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		// Read message from client
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		fmt.Printf("Received: %s\n", msg)

		// Echo message back to client
		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
