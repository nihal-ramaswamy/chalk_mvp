package websockets_impl

import (
	"database/sql"
	"sync"

	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
	"go.uber.org/zap"
)

type Room struct {
	room map[string]*ConferenceWsDto
	log  *zap.Logger
	mu   sync.Mutex
}

func NewRoom(log *zap.Logger) *Room {
	return &Room{
		room: make(map[string]*ConferenceWsDto),
		log:  log,
	}
}

func (r *Room) AddRoom(code string, db *sql.DB) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.room[code] = NewConferenceWsDto(r.log, db)
}

func (r *Room) DoesRoomExist(code string) bool {
	_, ok := r.room[code]
	return ok
}

func (r *Room) AddConnection(code string, connWithId *dto.ConnWithId) {
	r.room[code].AddConnection(connWithId)
}

func (r *Room) HandleWs(code string, connWithId *dto.ConnWithId) {
	r.room[code].HandleWs(connWithId)
}
