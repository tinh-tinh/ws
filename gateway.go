package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type Gateway struct {
	core.DynamicController
	module  core.Module
	upgrade websocket.Upgrader
}

func NewGateway(module core.Module) *Gateway {
	gateway := &Gateway{
		module: module,
		upgrade: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // allow all origins
			},
		},
	}

	return gateway
}
