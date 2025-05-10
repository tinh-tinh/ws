package ws

type Ctx interface{}

type DefaultCtx struct {
	messageType int
	data        []byte
	gateway     Gateway
}

func NewCtx(messageType int, data []byte, gateway Gateway) Ctx {
	return &DefaultCtx{
		messageType: messageType,
		data:        data,
		gateway:     gateway,
	}
}
