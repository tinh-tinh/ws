package ws

import (
	"github.com/gorilla/websocket"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

type Config struct {
	Path         string
	Upgrade      websocket.Upgrader
	Serializer   core.Encode
	Deserializer core.Decode
	Logger       *logger.Logger
	events       []*EventFnc
}

type Gateway struct {
	Config
	module core.Module
}

func NewGateway(module core.Module, config Config) *Gateway {
	gateway := &Gateway{
		module: module,
		Config: config,
	}

	return gateway
}
