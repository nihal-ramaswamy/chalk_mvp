package bookmark_api

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/interfaces"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type AuthGroup struct {
	routeHandlers []interfaces.HandlerInterface
	middlewares   []gin.HandlerFunc
}

func (*AuthGroup) Group() string {
	return "/bookmark"
}

func (h *AuthGroup) RouteHandlers() []interfaces.HandlerInterface {
	return h.routeHandlers
}

func NewAuthGroup(db *sql.DB, rdb *redis.Client, ctx context.Context, log *zap.Logger) *AuthGroup {
	handlers := []interfaces.HandlerInterface{}

	return &AuthGroup{
		routeHandlers: handlers,
		middlewares:   []gin.HandlerFunc{},
	}
}

func (*AuthGroup) AuthRequired() bool {
	return false
}

func (a *AuthGroup) Middlewares() []gin.HandlerFunc {
	return a.middlewares
}
