package ws

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	user     string
	room     string
	Messages []Message
}

func NewClient(url string, protocal string, origin string) *Client {
	conn, err := websocket.Dial(url, protocal, origin)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}

	return &Client{Conn: conn, Messages: []Message{}}
}

func (client *Client) Init(user string) *Client {
	client.user = user
	if err := websocket.Message.Send(client.Conn, user); err != nil {
		log.Fatal("Failed to send username:", err)
	}

	return client
}

func (client *Client) JoinRoom(room string) *Client {
	client.room = room
	if err := websocket.Message.Send(client.Conn, room); err != nil {
		log.Fatal("Failed to send room:", err)
	}
	return client
}

func (client *Client) Send(msg interface{}) {
	if err := websocket.JSON.Send(client.Conn, msg); err != nil {
		log.Println("Error sending message:", err)
		return
	}
}

func (client *Client) Lisen() {
	for {
		var msg Message
		if err := websocket.JSON.Receive(client.Conn, &msg); err != nil {
			log.Println("Connection closed:", err)
			return
		}
		// client.Messages = append(client.Messages, msg)
		fmt.Printf("\n[%s]: %s\n", msg.User, msg.Text)
	}
}
