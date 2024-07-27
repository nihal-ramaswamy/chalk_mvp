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

type SkillsHandler struct {
	db          *sql.DB
	log         *zap.Logger
	middlewares []gin.HandlerFunc
}

func NewSkillsHandler(
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

func (c *SkillsHandler) Pattern() string {
	return "/:category"
}

func (c *SkillsHandler) RequestMethod() string {
	return constants.GET
}

func (c *SkillsHandler) Middlewares() []gin.HandlerFunc {
	return c.middlewares
}

func (c *SkillsHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		category := ctx.Param("category")
		skills, err := db.GetValuesForCategory(c.db, category)

		if nil != err {
			utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"skills": skills,
		})
	}
}
