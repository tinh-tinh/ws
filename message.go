package ws

type Message struct {
	Event   string `json:"event"`
	Payload any    `json:"payload"`
}
