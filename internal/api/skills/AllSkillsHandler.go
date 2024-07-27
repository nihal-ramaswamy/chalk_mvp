package skills_api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/constants"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/db"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/utils"
	"go.uber.org/zap"
)

type AllSkillsHandler struct {
	db          *sql.DB
	log         *zap.Logger
	middlewares []gin.HandlerFunc
}

func NewAllSkillsHandler(
	db *sql.DB,
	log *zap.Logger,
	ctx context.Context,
) *AllSkillsHandler {
	return &AllSkillsHandler{
		db:          db,
		log:         log,
		middlewares: []gin.HandlerFunc{},
	}
}

func (c *AllSkillsHandler) Pattern() string {
	return "/"
}

func (c *AllSkillsHandler) RequestMethod() string {
	return constants.GET
}

func (c *AllSkillsHandler) Middlewares() []gin.HandlerFunc {
	return c.middlewares
}

func (c *AllSkillsHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		skills, err := db.GetAllValues(c.db)

		if nil != err {
			utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"skills": skills,
		})
	}
}
