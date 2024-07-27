package chat_api

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/interfaces"
	auth_middleware "github.com/nihal-ramaswamy/chalk_mvp/internal/middleware/auth"
	websockets_impl "github.com/nihal-ramaswamy/chalk_mvp/internal/websockets"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type ChatApi struct {
	middlewares   []gin.HandlerFunc
	routeHandlers []interfaces.HandlerInterface
}

func NewChatApi(
	pdb *sql.DB,
	rdb_auth *redis.Client,
	ctx context.Context,
	log *zap.Logger,
	upgrader *websocket.Upgrader,
	rdb_ws *redis.Client,
	roomDto *websockets_impl.Room,
) *ChatApi {
	handlers := []interfaces.HandlerInterface{
		NewChatHandler(upgrader, log, rdb_ws, rdb_auth, ctx, roomDto, pdb),
		NewChatViewHandler(pdb, log, rdb_auth, ctx),
	}

	return &ChatApi{
		middlewares:   []gin.HandlerFunc{auth_middleware.AuthMiddleware(pdb, rdb_auth, ctx, log)},
		routeHandlers: handlers,
	}
}

func (*ChatApi) Group() string {
	return "/chat"
}

func (c *ChatApi) Middlewares() []gin.HandlerFunc {
	return c.middlewares
}

func (c *ChatApi) RouteHandlers() []interfaces.HandlerInterface {
	return c.routeHandlers
}
