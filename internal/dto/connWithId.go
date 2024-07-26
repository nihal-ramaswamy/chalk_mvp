package dto

import "github.com/gorilla/websocket"

type ConnWithId struct {
	Conn *websocket.Conn
	ID   string
}

func NewConnWithId(conn *websocket.Conn, id string) *ConnWithId {
	return &ConnWithId{
		Conn: conn,
		ID:   id,
	}
}
