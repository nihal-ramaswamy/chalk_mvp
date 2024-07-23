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
		user := dto.NewStudent()
		if err := c.ShouldBindJSON(&user); nil != err {
			utils.HandleErrorAndAbortWithError(c, err, l.log, http.StatusBadRequest)
			return
		}
		if !db.DoesEmailExist(l.db, user.Email) {
			utils.HandleErrorAndAbortWithError(c, fmt.Errorf("user with email %s does not exist", user.Email), l.log, http.StatusUnauthorized)
			return
		}

		if !db.DoesPasswordMatch(l.db, user, l.log) {
			utils.HandleErrorAndAbortWithError(c, fmt.Errorf("invalid credentials"), l.log, http.StatusUnauthorized)
			return
		}

		token, err := utils.GenerateToken(user)
		if nil != err {
			utils.HandleErrorAndAbortWithError(c, err, l.log, http.StatusInternalServerError)
			return
		}

		l.rdb.Set(l.ctx, user.Email, token, constants.TOKEN_EXPIRY_TIME)

		c.JSON(http.StatusAccepted, gin.H{"token": token})
	}
}

func (l *LoginUserHandler) Middlewares() []gin.HandlerFunc {
	return l.middlewares
}
