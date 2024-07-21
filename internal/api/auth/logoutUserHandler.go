package auth_api

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

type LogoutUserHandler struct {
	ctx         context.Context
	rdb         *redis.Client
	log         *zap.Logger
	middlewares []gin.HandlerFunc
}

func NewLogoutUserHandler(pdb *sql.DB, rdb *redis.Client, ctx context.Context, log *zap.Logger) *LogoutUserHandler {
	return &LogoutUserHandler{
		rdb:         rdb,
		ctx:         ctx,
		log:         log,
		middlewares: []gin.HandlerFunc{auth_middleware.AuthMiddleware(pdb, rdb, ctx, log)},
	}
}

func (*LogoutUserHandler) Pattern() string {
	return "/signout"
}

func (*LogoutUserHandler) RequestMethod() string {
	return constants.POST
}

func (l *LogoutUserHandler) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetString("email")
		_, err := l.rdb.Del(l.ctx, email).Result()
		if err != nil {
			l.log.Error("Error deleting token from rdb")
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "ok"})
	}
}

func (l *LogoutUserHandler) Middlewares() []gin.HandlerFunc {
	return l.middlewares
}
