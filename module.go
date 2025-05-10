package ws

import (
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

const GATEWAY core.Provide = "WS_GATEWAY"

func Registry(configs ...Config) core.Modules {
	return func(module core.Module) core.Module {
		config := ParseConfig(configs...)
		sub := module.New(core.NewModuleOptions{})
		sub.NewProvider(core.ProviderOptions{
			Name:  GATEWAY,
			Value: &config,
		})
		sub.Export(GATEWAY)

		sub.Controllers(initHandler)
		return sub
	}
}
