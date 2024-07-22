package routes

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	auth_api "github.com/nihal-ramaswamy/chalk_mvp/internal/api/auth"
	bookmark_api "github.com/nihal-ramaswamy/chalk_mvp/internal/api/bookmark"
	healthcheck_api "github.com/nihal-ramaswamy/chalk_mvp/internal/api/healthcheck"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/constants"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/interfaces"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func NewRoutes(
	server *gin.Engine,
	log *zap.Logger,
	db *sql.DB,
	rdb_auth *redis.Client,
	ctx context.Context,
	rdb_ws *redis.Client,
) {
	serverGroupHandlers := []interfaces.ServerGroupInterface{
		healthcheck_api.NewHealthCheckGroup(db, rdb_auth, ctx, log),
		auth_api.NewAuthGroup(db, rdb_auth, ctx, log),
		bookmark_api.NewBookmarkGroup(db, rdb_auth, ctx, log),
	}

	for _, serverGroupHandler := range serverGroupHandlers {
		newGroup(server, serverGroupHandler)
	}
}

func newGroup(server *gin.Engine, groupHandler interfaces.ServerGroupInterface) {
	group := server.Group(groupHandler.Group(), groupHandler.Middlewares()...)
	{
		for _, route := range groupHandler.RouteHandlers() {
			newRoute(group, route)
		}
	}
}

func newRoute(server *gin.RouterGroup, routeHandler interfaces.HandlerInterface) {
	middlewares := routeHandler.Middlewares()
	middlewares = append(middlewares, routeHandler.Handler())
	switch routeHandler.RequestMethod() {
	case constants.GET:
		server.GET(routeHandler.Pattern(), middlewares...)
	case constants.POST:
		server.POST(routeHandler.Pattern(), middlewares...)
	}
}
