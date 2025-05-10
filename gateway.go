package ws

import (
	"encoding/json"
	"net/http"
	"reflect"

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
	ErrorHandler ErrorHandler
	events       []*EventFnc
}

func DefaultConfig() Config {
	logger := logger.Create(logger.Options{})
	return Config{
		Path: "",
		Upgrade: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		Serializer:   json.Marshal,
		Deserializer: json.Unmarshal,
		ErrorHandler: DefaultErrorHandler(logger),
		Logger:       logger,
	}
}

func ParseConfig(configs ...Config) Config {
	defaultConfig := DefaultConfig()
	for _, cfg := range configs {
		if cfg.Path != "" {
			defaultConfig.Path = cfg.Path
		}

		if !reflect.ValueOf(cfg).IsZero() {
			defaultConfig.Upgrade = cfg.Upgrade
		}

		if cfg.Serializer != nil {
			defaultConfig.Serializer = cfg.Serializer
		}

		if cfg.Deserializer != nil {
			defaultConfig.Deserializer = cfg.Deserializer
		}

		if cfg.ErrorHandler != nil {
			defaultConfig.ErrorHandler = cfg.ErrorHandler
		}

		if cfg.Logger != nil {
			defaultConfig.Logger = cfg.Logger
			if cfg.ErrorHandler == nil {
				defaultConfig.ErrorHandler = DefaultErrorHandler(cfg.Logger)
			}
		}
	}

	return defaultConfig
}
