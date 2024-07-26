package dto

import "time"

type Message struct {
	SentAt        time.Time `json:"sent_at,omitempty"`
	SenderId      string    `json:"sender,omitempty"`
	ReceiverId    string    `json:"receiver,omitempty"`
	Message       string    `json:"message"`
	ReceiverEmail string    `json:"receiver_email,omitempty"`
}

func NewMessage() *Message {
	return &Message{}
}
