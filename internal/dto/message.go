package dto

import "time"

type Message struct {
	SentAt        time.Time `json:"sent_at"`
	SenderEmail   string    `json:"sender"`
	ReceiverEmail string    `json:"receiver"`
}

func NewMessage() *Message {
	return &Message{}
}
