package dto

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type ConferenceWsDto struct {
	People map[*websocket.Conn]bool
	log    *zap.Logger
}

func NewConferenceWsDto(log *zap.Logger) *ConferenceWsDto {
	return &ConferenceWsDto{
		People: make(map[*websocket.Conn]bool),
		log:    log,
	}
}

func (c *ConferenceWsDto) AddConnection(ws *websocket.Conn) {
	c.People[ws] = true // TODO: Make this concurrent safe
}

func (c *ConferenceWsDto) RemoveConnection(ws *websocket.Conn) {
	c.People[ws] = false
	ws.Close()
}

func (c *ConferenceWsDto) HandleWs(ws *websocket.Conn) {
	c.readLoop(ws)
}

func (c *ConferenceWsDto) readLoop(ws *websocket.Conn) {
	var message Message

	for {
		err := ws.ReadJSON(&message)
		if nil != err {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				c.RemoveConnection(ws)
				break
			}
			c.log.Error("Error in connection", zap.Error(err))
			break
		}

		c.broadcast(message)
	}
}

func (c *ConferenceWsDto) broadcast(message Message) {
	for ws, active := range c.People {
		if !active {
			continue
		}
		go func(ws *websocket.Conn) {
			err := ws.WriteJSON(message)
			if nil != err {
				c.log.Error("error writing json", zap.Error(err))
			}
		}(ws)
	}
}
