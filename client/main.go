package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// Connect to the WebSocket server
	url := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	// Send a message to the server
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello from Go client!"))
	if err != nil {
		log.Println("Write error:", err)
		return
	}

	// Read a message from the server
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("Read error:", err)
		return
	}
	fmt.Printf("Received: %s\n", message)

	// Optional: Keep the connection open and listen for more messages
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}
			fmt.Println("Server:", string(msg))
		}
	}()

	// Send more messages in a loop
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		text := fmt.Sprintf("Ping %d", i)
		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}

	time.Sleep(2 * time.Second)
}
