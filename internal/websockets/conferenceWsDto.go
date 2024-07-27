package websockets_impl

import (
	"database/sql"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/db"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
	"go.uber.org/zap"
)

type ConferenceWsDto struct {
	db     *sql.DB
	People map[*dto.ConnWithId]bool
	log    *zap.Logger
	mu     sync.Mutex
}

func NewConferenceWsDto(log *zap.Logger, db *sql.DB) *ConferenceWsDto {
	return &ConferenceWsDto{
		db:     db,
		People: make(map[*dto.ConnWithId]bool),
		log:    log,
	}
}

func (c *ConferenceWsDto) AddConnection(connWithId *dto.ConnWithId) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.People[connWithId] = true
}

func (c *ConferenceWsDto) RemoveConnection(connWithId *dto.ConnWithId) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.People[connWithId] = false
	connWithId.Conn.Close()
}

func (c *ConferenceWsDto) HandleWs(connWithId *dto.ConnWithId, code string) {
	c.readLoop(connWithId, code)
}

func (c *ConferenceWsDto) readLoop(connWithId *dto.ConnWithId, code string) {
	var message dto.Message

	for {
		err := connWithId.Conn.ReadJSON(&message)
		if nil != err {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				c.RemoveConnection(connWithId)
				break
			}
			c.log.Error("Error in connection", zap.Error(err))
			break
		}

		message.SenderId = connWithId.ID
		message.ChatCode = code

		err = db.SaveChatToDb(c.db, message)

		if nil != err {
			c.log.Error("error saving chat", zap.Error(err))
		}

		c.broadcast(message)
	}
}

func (c *ConferenceWsDto) broadcast(message dto.Message) {
	for connWithId, active := range c.People {
		if !active {
			continue
		}
		go func(ws *websocket.Conn) {
			err := ws.WriteJSON(message)
			if nil != err {
				c.log.Error("error writing json", zap.Error(err))
			}
		}(connWithId.Conn)
	}
}
