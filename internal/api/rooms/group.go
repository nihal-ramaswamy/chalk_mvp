package rooms_api

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/interfaces"
	auth_middleware "github.com/nihal-ramaswamy/chalk_mvp/internal/middleware/auth"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RoomsApi struct {
	routeHandlers []interfaces.HandlerInterface
	middlewares   []gin.HandlerFunc
}

func (*RoomsApi) Group() string {
	return "/rooms"
}

func (h *RoomsApi) RouteHandlers() []interfaces.HandlerInterface {
	return h.routeHandlers
}

func NewRoomsApi(db *sql.DB, log *zap.Logger, rdb_auth *redis.Client, ctx context.Context, roomDto *dto.Room) *RoomsApi {
	handlers := []interfaces.HandlerInterface{
		NewCreateRoomHandler(db, log, roomDto),
	}

	return &RoomsApi{
		routeHandlers: handlers,
		middlewares:   []gin.HandlerFunc{auth_middleware.AuthMiddleware(db, rdb_auth, ctx, log)},
	}
}

func (*RoomsApi) AuthRequired() bool {
	return false
}

func (h *RoomsApi) Middlewares() []gin.HandlerFunc {
	return h.middlewares
}
