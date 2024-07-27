package skills_api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/interfaces"
	"go.uber.org/zap"
)

type SkillsApi struct {
	middlewares   []gin.HandlerFunc
	routeHandlers []interfaces.HandlerInterface
}

func NewChatApi(
	pdb *sql.DB,
	log *zap.Logger,
) *SkillsApi {
	handlers := []interfaces.HandlerInterface{
		NewAllSkillsHandler(pdb, log),
		NewSkillsHandler(pdb, log),
	}

	return &SkillsApi{
		middlewares:   []gin.HandlerFunc{},
		routeHandlers: handlers,
	}
}

func (*SkillsApi) Group() string {
	return "/skill"
}

func (c *SkillsApi) Middlewares() []gin.HandlerFunc {
	return c.middlewares
}

func (c *SkillsApi) RouteHandlers() []interfaces.HandlerInterface {
	return c.routeHandlers
}
