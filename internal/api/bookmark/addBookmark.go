package bookmark_api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/constants"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/db"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
	auth_middleware "github.com/nihal-ramaswamy/chalk_mvp/internal/middleware/auth"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type AddBookmarkHandler struct {
	db          *sql.DB
	log         *zap.Logger
	middlewares []gin.HandlerFunc
}

func NewAddBookmarkHandler(pdb *sql.DB, rdb *redis.Client, ctx context.Context, log *zap.Logger) *AddBookmarkHandler {
	return &AddBookmarkHandler{
		db:          pdb,
		log:         log,
		middlewares: []gin.HandlerFunc{auth_middleware.AuthMiddleware(pdb, rdb, ctx, log)},
	}
}

func (*AddBookmarkHandler) Pattern() string {
	return "/add"
}

func (n *AddBookmarkHandler) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetString("email")
		addBookmarkStruct := dto.NewBookmark()

		if err := c.ShouldBindJSON(&addBookmarkStruct); nil != err {
			err := c.Error(err)
			n.log.Info("Responding with error", zap.Error(err))
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		n.log.Info("Debug: ", zap.String("email", email), zap.Any("bookmark ", addBookmarkStruct))

		if err := db.AppendEmailToBookmarks(n.db, email, addBookmarkStruct); nil != err {
			err := c.Error(err)
			n.log.Info("Responding with error", zap.Error(err))
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "ok"})
	}
}

func (*AddBookmarkHandler) RequestMethod() string {
	return constants.POST
}

func (n *AddBookmarkHandler) Middlewares() []gin.HandlerFunc {
	return n.middlewares
}
