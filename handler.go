package ws

import (
	"fmt"
	"net/http"

	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func initHandler(module core.Module) core.Controller {
	gateway, ok := module.Ref(GATEWAY).(*Config)
	if !ok {
		panic("not register gateway")
	}
	ctrl := module.NewController(gateway.Path)

	ctrl.Handler("", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your handler logic here
		conn, err := gateway.Upgrade.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Upgrade error:", err)
			return
		}
		defer conn.Close()

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
	}))

	return ctrl
}
