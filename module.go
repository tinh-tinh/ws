package ws

import (
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"golang.org/x/net/websocket"
)

type Handler interface {
	OnConnect(conn *websocket.Conn)
}

func handler(module core.Module) core.Controller {
	socketConfig, ok := module.Ref(WEBSOCKET).(*Options)
	if !ok {
		panic("not have config websocket server")
	}

	ctrl := module.NewController(socketConfig.Prefix)

	if socketConfig.Handler != nil {
		ctrl.Handler("", websocket.Handler(socketConfig.Handler.OnConnect))
	} else {
		ctrl.Handler("", websocket.Handler(DefaultHandler))
	}

	return ctrl
}

type Options struct {
	Prefix  string
	Handler Handler
}

const WEBSOCKET core.Provide = "WEBSOCKET"

func Register(opt Options) core.Modules {
	return func(module core.Module) core.Module {
		wsModule := module.New(core.NewModuleOptions{})

		wsModule.NewProvider(core.ProviderOptions{
			Name:  WEBSOCKET,
			Value: &opt,
		})
		wsModule.Export(WEBSOCKET)
		wsModule.Controllers(handler)

		return wsModule
	}
}
