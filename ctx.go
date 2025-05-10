package ws

type Ctx interface {
	GetMessage() Message
	Send(v any) (*Message, error)
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

func (c *DefaultCtx) Send(v any) (*Message, error) {
	msg := c.message
	msg.Payload = v
	return &msg, nil
}
