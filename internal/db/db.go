package db

import (
	"database/sql"

	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func CreateNewMeeting(db *sql.DB, conference *dto.Conference) error {
	return insertIntoConference(db, conference)
}

func DoesEmailExist(db *sql.DB, email string) bool {
	_, err := selectAllFromUserWhereEmailIs(db, email)
	return err != sql.ErrNoRows
}

func RegisterNewUser(db *sql.DB, user *dto.Student, log *zap.Logger) string {
	user = user.HashAndSalt()

	id, err := insertIntoUser(db, user)
	if err != nil {
		log.Error(err.Error())
	}

	return id
}

func DoesPasswordMatch(db *sql.DB, user *dto.Student, log *zap.Logger) bool {
	password, err := selectPasswordFromUserWhereEmailIDs(db, user.Email)

	if nil != err {
		log.Error(err.Error())
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)) == nil
}

func GetUserFromEmail(db *sql.DB, email string) (dto.Student, error) {
	return selectAllFromUserWhereEmailIs(db, email)
}
