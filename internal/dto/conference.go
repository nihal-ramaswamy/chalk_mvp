package dto

import "time"

type Conference struct {
	CreatedAt time.Time `json:"created_at"`
	Code      string    `json:"code"`
	Admin     string    `json:"admin"`
	Active    bool      `json:"active"`
}

func NewConference(code, admin string) *Conference {
	return &Conference{
		Code:   code,
		Admin:  admin,
		Active: true,
	}
}
