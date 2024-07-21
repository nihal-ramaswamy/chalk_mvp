package dto

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Room struct {
	room map[string]*ConferenceWsDto
	log  *zap.Logger
}

func NewRoom(log *zap.Logger) *Room {
	return &Room{
		room: make(map[string]*ConferenceWsDto),
		log:  log,
	}
}

func (r *Room) AddRoom(code string) {
	r.room[code] = NewConferenceWsDto(r.log)
}

func (r *Room) DoesRoomExist(code string) bool {
	_, ok := r.room[code]
	return ok
}

func (r *Room) AddConnection(code string, ws *websocket.Conn) {
	r.room[code].AddConnection(ws)
}

func (r *Room) HandleWs(code string, ws *websocket.Conn) {
	r.room[code].HandleWs(ws)
}
