package bookmark_api

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/interfaces"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type BookmarkGroup struct {
	routeHandlers []interfaces.HandlerInterface
	middlewares   []gin.HandlerFunc
}

func (*BookmarkGroup) Group() string {
	return "/bookmark"
}

func (h *BookmarkGroup) RouteHandlers() []interfaces.HandlerInterface {
	return h.routeHandlers
}

func NewBookmarkGroup(db *sql.DB, rdb *redis.Client, ctx context.Context, log *zap.Logger) *BookmarkGroup {
	handlers := []interfaces.HandlerInterface{
		NewViewBookmarkHandler(db, rdb, ctx, log),
		NewAddBookmarkHandler(db, rdb, ctx, log),
	}

	return &BookmarkGroup{
		routeHandlers: handlers,
		middlewares:   []gin.HandlerFunc{},
	}
}

func (*BookmarkGroup) AuthRequired() bool {
	return true
}

func (a *BookmarkGroup) Middlewares() []gin.HandlerFunc {
	return a.middlewares
}
