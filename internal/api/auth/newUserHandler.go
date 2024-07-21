package auth_api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/constants"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/db"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
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
			err := c.Error(err)
			n.log.Info("Responding with error", zap.Error(err))
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		n.log.Info("Student: ", zap.Any("student", student))

		if db.DoesEmailExist(n.db, student.Email) {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("User with email %s already exists", student.Email)})
			return
		}

		id := db.RegisterNewUser(n.db, student, n.log)

		c.JSON(http.StatusAccepted, gin.H{"id": id})
	}
}

func (*NewStudentHandler) RequestMethod() string {
	return constants.POST
}

func (n *NewStudentHandler) Middlewares() []gin.HandlerFunc {
	return n.middlewares
}
