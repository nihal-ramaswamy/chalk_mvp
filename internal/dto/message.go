package dto

import "time"

type Message struct {
	SentAt        time.Time `json:"sent_at,omitempty"`
	SenderEmail   string    `json:"sender,omitempty"`
	ReceiverEmail string    `json:"receiver"`
}

func NewMessage() *Message {
	return &Message{}
}
