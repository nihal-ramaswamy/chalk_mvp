package skills_api

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/interfaces"
	websockets_impl "github.com/nihal-ramaswamy/chalk_mvp/internal/websockets"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type SkillsApi struct {
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
) *SkillsApi {
	handlers := []interfaces.HandlerInterface{}

	return &SkillsApi{
		middlewares:   []gin.HandlerFunc{},
		routeHandlers: handlers,
	}
}

func (*SkillsApi) Group() string {
	return "/skill"
}

func (c *SkillsApi) Middlewares() []gin.HandlerFunc {
	return c.middlewares
}

func (c *SkillsApi) RouteHandlers() []interfaces.HandlerInterface {
	return c.routeHandlers
}
