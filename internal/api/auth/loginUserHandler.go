package auth_api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/constants"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/db"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/utils"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type LoginUserHandler struct {
	log         *zap.Logger
	db          *sql.DB
	rdb         *redis.Client
	ctx         context.Context
	middlewares []gin.HandlerFunc
}

func NewLoginUserHandler(db *sql.DB, rdb *redis.Client, ctx context.Context, log *zap.Logger) *LoginUserHandler {
	return &LoginUserHandler{
		log:         log,
		db:          db,
		rdb:         rdb,
		ctx:         ctx,
		middlewares: []gin.HandlerFunc{},
	}
}

func (*LoginUserHandler) Pattern() string {
	return "/signin"
}

func (*LoginUserHandler) RequestMethod() string {
	return constants.POST
}

func (l *LoginUserHandler) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := dto.NewUser()
		if err := c.ShouldBindJSON(&user); nil != err {
			err := c.Error(err)
			l.log.Info("Responding with error", zap.Error(err))
			c.AbortWithStatus(http.StatusBadRequest)
		}
		if !db.DoesEmailExist(l.db, user.Email) {
			c.JSON(http.StatusUnauthorized,
				gin.H{"error": fmt.Sprintf("User with email %s does not exist", user.Email)})
			return
		}

		if !db.DoesPasswordMatch(l.db, user, l.log) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		token, err := utils.GenerateToken(user)
		if nil != err {
			err := c.Error(err)
			l.log.Info("Responding with error", zap.Error(err))

			c.AbortWithStatus(http.StatusInternalServerError)
		}

		l.rdb.Set(l.ctx, user.Email, token, constants.TOKEN_EXPIRY_TIME)

		c.JSON(http.StatusAccepted, gin.H{"token": token})
	}
}

func (l *LoginUserHandler) Middlewares() []gin.HandlerFunc {
	return l.middlewares
}
