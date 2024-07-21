package routes

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	auth_api "github.com/nihal-ramaswamy/chalk_mvp/internal/api/auth"
	chat_api "github.com/nihal-ramaswamy/chalk_mvp/internal/api/chat"
	healthcheck_api "github.com/nihal-ramaswamy/chalk_mvp/internal/api/healthcheck"
	rooms_api "github.com/nihal-ramaswamy/chalk_mvp/internal/api/rooms"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/constants"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/interfaces"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func NewRoutes(
	server *gin.Engine,
	upgrader *websocket.Upgrader,
	log *zap.Logger,
	db *sql.DB,
	rdb_auth *redis.Client,
	ctx context.Context,
	rdb_ws *redis.Client,
	roomDto *dto.Room,
) {
	serverGroupHandlers := []interfaces.ServerGroupInterface{
		healthcheck_api.NewHealthCheckGroup(db, rdb_auth, ctx, log),
		auth_api.NewAuthGroup(db, rdb_auth, ctx, log),
		rooms_api.NewRoomsApi(db, log, rdb_auth, ctx, roomDto),
		chat_api.NewChatApi(db, rdb_auth, ctx, log, upgrader, rdb_ws, roomDto),
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
