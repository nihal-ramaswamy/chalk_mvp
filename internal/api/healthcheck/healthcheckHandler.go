package healthcheck_api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/constants"
	auth_middleware "github.com/nihal-ramaswamy/chalk_mvp/internal/middleware/auth"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type HealthCheckHandler struct {
	middlewares []gin.HandlerFunc
}

func NewHealthCheckHandler(pdb *sql.DB, rdb *redis.Client, ctx context.Context, log *zap.Logger) *HealthCheckHandler {
	return &HealthCheckHandler{
		middlewares: []gin.HandlerFunc{auth_middleware.AuthMiddleware(pdb, rdb, ctx, log)},
	}
}

func (*HealthCheckHandler) Pattern() string {
	return "/healthcheck"
}

func (*HealthCheckHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

func (*HealthCheckHandler) RequestMethod() string {
	return constants.GET
}

func (h *HealthCheckHandler) Middlewares() []gin.HandlerFunc {
	return h.middlewares
}
