package chat_api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/constants"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/db"
	auth_middleware "github.com/nihal-ramaswamy/chalk_mvp/internal/middleware/auth"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/utils"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type ChatViewHandler struct {
	db          *sql.DB
	log         *zap.Logger
	middlewares []gin.HandlerFunc
}

func NewChatViewHandler(
	db *sql.DB,
	log *zap.Logger,
	rdb_auth *redis.Client,
	ctx context.Context,
) *ChatViewHandler {
	return &ChatViewHandler{
		db:          db,
		log:         log,
		middlewares: []gin.HandlerFunc{auth_middleware.AuthMiddleware(db, rdb_auth, ctx, log)},
	}
}

func (c *ChatViewHandler) Pattern() string {
	return "/chat/:code"
}

func (c *ChatViewHandler) RequestMethod() string {
	return constants.GET
}

func (c *ChatViewHandler) Middlewares() []gin.HandlerFunc {
	return c.middlewares
}

func (c *ChatViewHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString("email")
		code := ctx.Param("code")

		id, err := db.GetStudentIdFromEmail(c.db, email)

		if nil != err {
			utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)
			return
		}
		ok, err := db.IsStudentAuthorizedToViewChat(c.db, id, code)
		if nil != err {
			utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)
			return
		}

		if !ok {
			utils.HandleErrorAndAbortWithError(ctx, fmt.Errorf("unauthorized to view chat"), c.log, http.StatusUnauthorized)
			return
		}

		messages, err := db.ViewChat(c.db, code)

		if nil != err {
			utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"chat": messages,
		})
	}
}
