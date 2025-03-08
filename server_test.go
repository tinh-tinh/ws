package ws_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/tinh-tinh/ws"
	"golang.org/x/net/websocket"
)

func TestWebsocket(t *testing.T) {
	http.Handle("/ws", websocket.Handler(ws.DefaultHandler))
	log.Println("Chat server started on :8080")
	go http.ListenAndServe(":8080", nil)

	conn, err := websocket.Dial("ws://localhost:8080/ws", "", "http://localhost/")
	if err != nil {
		log.Fatal("Failed to connect to WebSocket server:", err)
	}
	defer conn.Close()
}
