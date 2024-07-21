package rooms_api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/constants"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/db"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/utils"
	"go.uber.org/zap"
)

type CreateRoomHandler struct {
	db          *sql.DB
	log         *zap.Logger
	roomDto     *dto.Room
	middlewares []gin.HandlerFunc
}

func NewCreateRoomHandler(db *sql.DB, log *zap.Logger, roomDto *dto.Room) *CreateRoomHandler {
	return &CreateRoomHandler{
		db:          db,
		log:         log,
		roomDto:     roomDto,
		middlewares: []gin.HandlerFunc{},
	}
}

func (*CreateRoomHandler) Pattern() string {
	return "/create"
}

func (*CreateRoomHandler) RequestMethod() string {
	return constants.POST
}

func (c *CreateRoomHandler) Middlewares() []gin.HandlerFunc {
	return c.middlewares
}

func (c *CreateRoomHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString("email")

		code := utils.NewCode(9)
		conference := dto.NewConference(code, email)
		err := db.CreateNewMeeting(c.db, conference)

		c.roomDto.AddRoom(code)

		utils.HandleErrorAndAbortWithError(ctx, err, c.log)

		ctx.JSON(http.StatusOK, gin.H{
			"code": code,
		})
	}
}
