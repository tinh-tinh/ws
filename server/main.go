package main

import (
	"log"
	"net/http"

	"github.com/tinh-tinh/ws"
	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/ws", websocket.Handler(ws.Handler))
	log.Println("Chat server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
