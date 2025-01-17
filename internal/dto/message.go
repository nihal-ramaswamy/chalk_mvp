package dto

import "time"

type Message struct {
	Id       string    `json:"id,omitempty"`
	SentAt   time.Time `json:"sent_at,omitempty"`
	SenderId string    `json:"sender,omitempty"`
	ChatCode string    `json:"chat_code,omitempty"`
	Message  string    `json:"message"`
}

func NewMessage() *Message {
	return &Message{}
}
