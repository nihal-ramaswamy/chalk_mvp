package healthcheck_api

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/interfaces"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type HealthCheckGroup struct {
	routeHandlers []interfaces.HandlerInterface
	middlewares   []gin.HandlerFunc
}

func (*HealthCheckGroup) Group() string {
	return "/healthcheck"
}

func (h *HealthCheckGroup) RouteHandlers() []interfaces.HandlerInterface {
	return h.routeHandlers
}

func NewHealthCheckGroup(pdb *sql.DB, rdb *redis.Client, ctx context.Context, log *zap.Logger) *HealthCheckGroup {
	handlers := []interfaces.HandlerInterface{
		NewHealthCheckHandler(pdb, rdb, ctx, log),
	}

	return &HealthCheckGroup{
		routeHandlers: handlers,
		middlewares:   []gin.HandlerFunc{},
	}
}

func (*HealthCheckGroup) AuthRequired() bool {
	return false
}

func (h *HealthCheckGroup) Middlewares() []gin.HandlerFunc {
	return h.middlewares
}
