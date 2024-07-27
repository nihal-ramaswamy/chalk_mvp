package chat_api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/constants"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/db"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
	auth_middleware "github.com/nihal-ramaswamy/chalk_mvp/internal/middleware/auth"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/utils"
	websockets_impl "github.com/nihal-ramaswamy/chalk_mvp/internal/websockets"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type ChatWsHandler struct {
	log         *zap.Logger
	upgrader    *websocket.Upgrader
	rdb         *redis.Client
	ctx         context.Context
	roomDto     *websockets_impl.Room
	db          *sql.DB
	middlewares []gin.HandlerFunc
}

func NewChatHandler(
	upgrader *websocket.Upgrader,
	log *zap.Logger,
	rdb_ws *redis.Client,
	rdb_auth *redis.Client,
	ctx context.Context,
	roomDto *websockets_impl.Room,
	db *sql.DB,
) *ChatWsHandler {
	return &ChatWsHandler{
		log:         log,
		upgrader:    upgrader,
		rdb:         rdb_ws,
		ctx:         ctx,
		db:          db,
		roomDto:     roomDto,
		middlewares: []gin.HandlerFunc{auth_middleware.AuthMiddleware(db, rdb_auth, ctx, log)},
	}
}

func (*ChatWsHandler) Pattern() string {
	return "/ws/:code"
}

func (*ChatWsHandler) RequestMethod() string {
	return constants.GET
}

func (c *ChatWsHandler) Middlewares() []gin.HandlerFunc {
	return c.middlewares
}

func (c *ChatWsHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Param("code")
		email := ctx.GetString("email")

		id, err := db.GetStudentIdFromEmail(c.db, email)
		if nil != err {
			utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)
			return
		}
		conn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if nil != err {
			utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)
			return
		}
		connWithId := dto.NewConnWithId(conn, id)

		defer func() {
			err := conn.Close()
			if nil != err {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
					c.log.Info("closing connection", zap.String("id", connWithId.ID))
					return
				}
				c.log.Error("error closing connection", zap.Error(err))
				utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)
				return
			}
		}()

		c.roomDto.AddConnection(code, connWithId)
		c.roomDto.HandleWs(code, connWithId)
	}
}
