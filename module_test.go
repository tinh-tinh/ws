package ws_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/ws"
)

func Test_Ws(t *testing.T) {
	eventGateway := func(module core.Module) core.Provider {
		handler := ws.NewHandler(module)

		handler.SubscribeMessage("on-ack", func(ctx ws.Ctx) error {
			fmt.Println(ctx.GetMessage())
			return nil
		})

		return handler
	}

	eventModule := func() core.Module {
		module := core.NewModule(core.NewModuleOptions{
			Imports: []core.Modules{ws.Registry(ws.Config{
				Path: "/ws",
				Upgrade: websocket.Upgrader{
					CheckOrigin: func(r *http.Request) bool {
						return true
					},
				},
			})},
			Providers: []core.Providers{eventGateway},
		})

		return module
	}

	app := core.CreateFactory(eventModule)

	testServer := httptest.NewServer(app.PrepareBeforeListen())
	defer testServer.Close()

	wsUrl := strings.ReplaceAll(testServer.URL, "http", "ws") + "/ws/"
	conn, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	// Send a message to the server
	err = conn.WriteMessage(websocket.TextMessage, []byte(`{"event":"on-ack", "payload": "abc"}`))
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
}
