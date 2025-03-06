package ws

import (
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"golang.org/x/net/websocket"
)

func handler(module core.Module) core.Controller {
	ctrl := module.NewController("ws")

	ctrl.Handler("", websocket.Handler(Handler))

	return ctrl
}

func Register() core.Modules {
	return func(module core.Module) core.Module {
		wsModule := module.New(core.NewModuleOptions{})

		wsModule.Controllers(handler)

		return wsModule
	}
}
