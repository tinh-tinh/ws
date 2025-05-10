package ws

import (
	"net/http"
	"slices"

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
			gateway.ErrorHandler(err)
			return
		}
		defer conn.Close()

		for {
			// Read message from client
			messageType, msg, err := conn.ReadMessage()
			if err != nil {
				gateway.ErrorHandler(err)
				break
			}

			// Parse data
			var message Message
			err = gateway.Deserializer(msg, &message)
			if err != nil {
				gateway.ErrorHandler(err)
				break
			}

			// Find Event
			subscriberIdx := slices.IndexFunc(gateway.events, func(e *EventFnc) bool {
				return e.Event == message.Event
			})
			if subscriberIdx == -1 {
				continue
			}

			// Handler
			ctx := NewCtx(messageType, message, *gateway)
			res, err := gateway.events[subscriberIdx].Handler(ctx)
			if err != nil {
				gateway.ErrorHandler(err)
				break
			}

			if res != nil {
				msg, err := gateway.Serializer(res)
				if err != nil {
					gateway.ErrorHandler(err)
					break
				}
				// Echo message back to client
				err = conn.WriteMessage(messageType, msg)
				if err != nil {
					gateway.ErrorHandler(err)
					break
				}
			}
		}
	}))

	return ctrl
}
