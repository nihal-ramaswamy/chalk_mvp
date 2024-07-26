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
	"github.com/nihal-ramaswamy/chalk_mvp/internal/utils"
	websockets_impl "github.com/nihal-ramaswamy/chalk_mvp/internal/websockets"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type ChatHandler struct {
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
	ctx context.Context,
	roomDto *websockets_impl.Room,
	db *sql.DB,
) *ChatHandler {
	return &ChatHandler{
		log:      log,
		upgrader: upgrader,
		rdb:      rdb_ws,
		ctx:      ctx,
		db:       db,
		roomDto:  roomDto,
	}
}

func (*ChatHandler) Pattern() string {
	return "/:code"
}

func (*ChatHandler) RequestMethod() string {
	return constants.GET
}

func (c *ChatHandler) Middlewares() []gin.HandlerFunc {
	return c.middlewares
}

func (c *ChatHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Param("code")
		email := ctx.GetString("email")

		id, err := db.GetStudentIdFromEmail(c.db, email)
		utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)

		conn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)

		connWithId := dto.NewConnWithId(conn, id)

		defer func() {
			err := conn.Close()
			if nil != err {
				c.log.Error("error closing connection", zap.Error(err))
				utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)
				return
			}
		}()

		c.roomDto.AddConnection(code, connWithId)
		c.roomDto.HandleWs(code, connWithId)
	}
}
