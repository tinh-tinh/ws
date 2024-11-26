package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tinh-tinh/ws"
)

func main() {
	client := ws.NewClient("ws://localhost:8080/ws", "", "http://localhost/")
	defer client.Conn.Close()
	// Ask for username
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	// Ask for room to join
	fmt.Print("Enter the room you want to join: ")
	room, _ := reader.ReadString('\n')
	room = strings.TrimSpace(room)

	client.Init(username).JoinRoom(room)

	go client.Lisen()

	// Read user input and send messages
	for {
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Exiting chat...")
			break
		}

		msg := ws.Message{
			User: username,
			Text: text,
			Room: room,
		}

		client.Send(msg)
	}
}
