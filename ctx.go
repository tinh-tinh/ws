package ws

type Ctx interface {
	GetMessage() Message
}

type DefaultCtx struct {
	messageType int
	message     Message
	config      Config
}

func NewCtx(messageType int, message Message, config Config) Ctx {
	return &DefaultCtx{
		messageType: messageType,
		message:     message,
		config:      config,
	}
}

func (c *DefaultCtx) GetMessage() Message {
	return c.message
}
