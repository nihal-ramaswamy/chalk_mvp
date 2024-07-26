package rooms_api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/constants"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/db"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/utils"
	websockets_impl "github.com/nihal-ramaswamy/chalk_mvp/internal/websockets"
	"go.uber.org/zap"
)

type CreateRoomHandler struct {
	db          *sql.DB
	log         *zap.Logger
	roomDto     *websockets_impl.Room
	middlewares []gin.HandlerFunc
}

func NewCreateRoomHandler(db *sql.DB, log *zap.Logger, roomDto *websockets_impl.Room) *CreateRoomHandler {
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

		id1, err := db.GetStudentIdFromEmail(c.db, email)
		if nil != err {
			utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)
			return
		}

		otherStudent := dto.NewBookmark()
		if err := ctx.ShouldBindJSON(&otherStudent); nil != err {
			utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)
			return
		}

		id2, err := db.GetStudentIdFromEmail(c.db, otherStudent.StudentEmail)
		if nil != err {
			utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)
			return
		}

		code, err := db.GetCode(c.db, id1, id2)

		if nil != err {
			utils.HandleErrorAndAbortWithError(ctx, err, c.log, http.StatusInternalServerError)
			return
		}

		c.roomDto.AddRoom(code, c.db)

		ctx.JSON(http.StatusOK, gin.H{
			"code": code,
		})
	}
}
