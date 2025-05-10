package ws

import "github.com/tinh-tinh/tinhtinh/v2/core"

type HandleFnc func(ctx Ctx) error

type EventFnc struct {
	Event   string
	Handler HandleFnc
}

type Handler struct {
	core.DynamicProvider
	module core.Module
}

func NewHandler(module core.Module) *Handler {
	return &Handler{
		module: module,
	}
}

func (h *Handler) SubscribeMessage(event string, handler HandleFnc) {
	eventFnc := &EventFnc{
		Event:   event,
		Handler: handler,
	}

	core.InitProviders(h.module, core.ProviderOptions{
		Name: GetEventProvider(event),
		Factory: func(param ...interface{}) interface{} {
			config, ok := param[0].(*Config)
			if !ok {
				return nil
			}
			config.events = append(config.events, eventFnc)

			return config.events
		},
		Inject: []core.Provide{GATEWAY},
		Scope:  h.Scope,
	})
}

func GetEventProvider(event string) core.Provide {
	return core.Provide("WS_" + event)
}
