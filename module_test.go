package ws_test

import (
	"log"
	"testing"

	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/ws"
	"golang.org/x/net/websocket"
)

func Test_Module(t *testing.T) {
	appModule := func() core.Module {
		appM := core.NewModule(core.NewModuleOptions{
			Imports: []core.Modules{ws.Register(ws.Options{
				Prefix: "ws",
			})},
		})

		return appM
	}

	app := core.CreateFactory(appModule)

	go app.Listen(8000)
	conn, err := websocket.Dial("ws://localhost:8000/ws/", "", "http://localhost/")
	if err != nil {
		log.Fatal("Failed to connect to WebSocket server:", err)
	}
	defer conn.Close()
}
