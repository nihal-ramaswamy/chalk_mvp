package auth_api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/constants"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/db"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/utils"
	"go.uber.org/zap"
)

type NewStudentHandler struct {
	db          *sql.DB
	log         *zap.Logger
	middlewares []gin.HandlerFunc
}

func NewNewUserHandler(db *sql.DB, log *zap.Logger) *NewStudentHandler {
	return &NewStudentHandler{
		db:  db,
		log: log,
	}
}

func (*NewStudentHandler) Pattern() string {
	return "/register"
}

func (n *NewStudentHandler) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		student := dto.NewStudent()

		if err := c.ShouldBindJSON(&student); err != nil {
			utils.HandleErrorAndAbortWithError(c, err, n.log, http.StatusBadRequest)
			return
		}

		if db.DoesEmailExist(n.db, student.Email) {
			utils.HandleErrorAndAbortWithError(c, fmt.Errorf("user with email %s already exists", student.Email), n.log, http.StatusUnprocessableEntity)
			return
		}

		id := db.RegisterNewUser(n.db, student, n.log)

		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func (*NewStudentHandler) RequestMethod() string {
	return constants.POST
}

func (n *NewStudentHandler) Middlewares() []gin.HandlerFunc {
	return n.middlewares
}
