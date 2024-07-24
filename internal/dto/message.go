package dto

import "time"

type Message struct {
	SentAt        time.Time `json:"sent_at,omitempty"`
	SenderEmail   string    `json:"sender,omitempty"`
	ReceiverEmail string    `json:"receiver"`
	Message       string    `json:"message"`
}

func NewMessage() *Message {
	return &Message{}
}
