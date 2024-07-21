package dto

type Message struct {
	Message string `json:"message"`
	Person  string `json:"person"`
}

func NewMessage() *Message {
	return &Message{}
}
