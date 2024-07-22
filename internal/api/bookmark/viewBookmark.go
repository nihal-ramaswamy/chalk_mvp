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

type ViewBookmark struct {
	db          *sql.DB
	log         *zap.Logger
	middlewares []gin.HandlerFunc
}

func NewViewBookmarkHandler(pdb *sql.DB, rdb *redis.Client, ctx context.Context, log *zap.Logger) *ViewBookmark {
	return &ViewBookmark{
		db:          pdb,
		log:         log,
		middlewares: []gin.HandlerFunc{auth_middleware.AuthMiddleware(pdb, rdb, ctx, log)},
	}
}

func (*ViewBookmark) Pattern() string {
	return "/view"
}

func (n *ViewBookmark) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookmark := dto.NewBookmark()

		if err := c.ShouldBindJSON(&bookmark); nil != err {
			err := c.Error(err)
			n.log.Info("Responding with error", zap.Error(err))
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		bookmarks, err := db.GetBookmarksForUser(n.db, bookmark.StudentEmail)

		if nil != err {
			n.log.Error("Error fetching bookmarks", zap.String("email", bookmark.StudentEmail))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, bookmarks)
	}
}

func (*ViewBookmark) RequestMethod() string {
	return constants.POST
}

func (n *ViewBookmark) Middlewares() []gin.HandlerFunc {
	return n.middlewares
}
