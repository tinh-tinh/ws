package ws

import (
	"log"

	"golang.org/x/net/websocket"
)

type Message struct {
	User string `json:"user"`
	Text string `json:"text"`
	Room string `json:"room"`
}

type Room struct {
	Clients   map[*websocket.Conn]string
	Broadcast chan Message
}

var rooms = make(map[string]*Room) // Map of rooms by room name

func Handler(conn *websocket.Conn) {
	defer conn.Close()
	// Read the user and room
	var user, room string
	if err := websocket.Message.Receive(conn, &user); err != nil {
		log.Println("Failed to get user:", err)
		return
	}

	// Ask for the room to join
	if err := websocket.Message.Receive(conn, &room); err != nil {
		log.Println("Failed to get room:", err)
		return
	}

	// If the room doesn't exist, create it
	if _, ok := rooms[room]; !ok {
		rooms[room] = &Room{
			Clients:   make(map[*websocket.Conn]string),
			Broadcast: make(chan Message),
		}
		go HandlerRoom(room)
	}

	// Add client to the room
	rooms[room].Clients[conn] = user
	log.Printf("%s joined room: %s", user, room)

	// Listen for incoming messages from the client
	go func() {
		for {
			var msg Message
			if err := websocket.JSON.Receive(conn, &msg); err != nil {
				log.Printf("%s disconnected from room %s: %v", user, room, err)
				delete(rooms[room].Clients, conn)
				break
			}
			msg.User = user
			msg.Room = room
			rooms[room].Broadcast <- msg
		}
	}()

	// Receive messages and broadcast to all clients in the room
	for msg := range rooms[room].Broadcast {
		for client := range rooms[room].Clients {
			if client != conn {
				if err := websocket.JSON.Send(client, msg); err != nil {
					log.Println("Error sending message:", err)
					client.Close()
					delete(rooms[room].Clients, client)
				}
			}
		}
	}
}

// Handle message broadcasting in a room
func HandlerRoom(room string) {
	for {
		msg := <-rooms[room].Broadcast
		for client := range rooms[room].Clients {
			if err := websocket.JSON.Send(client, msg); err != nil {
				log.Println("Error sending message:", err)
				client.Close()
				delete(rooms[room].Clients, client)
			}
		}
	}
}
