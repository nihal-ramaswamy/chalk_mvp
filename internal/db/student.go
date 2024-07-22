package db

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func DoesEmailExist(db *sql.DB, email string) bool {
	_, err := selectAllFromStudentWhereEmailIs(db, email)
	return err != sql.ErrNoRows
}

func RegisterNewUser(db *sql.DB, user *dto.Student, log *zap.Logger) string {
	user = user.HashAndSalt()

	id, err := insertIntoStudent(db, user)
	if err != nil {
		log.Error(err.Error())
	}

	return id
}

func DoesPasswordMatch(db *sql.DB, user *dto.Student, log *zap.Logger) bool {
	password, err := selectPasswordFromStudentWhereEmailIDs(db, user.Email)

	if nil != err {
		log.Error(err.Error())
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)) == nil
}

func GetUserFromEmail(db *sql.DB, email string) (dto.Student, error) {
	return selectAllFromStudentWhereEmailIs(db, email)
}

// -----Queries -----

func insertIntoStudent(db *sql.DB, user *dto.Student) (string, error) {
	if db == nil {
		panic("db cannot be nil")
	}

	var id string
	query := `INSERT INTO "STUDENT" (NAME, EMAIL, PASSWORD, DESCRIPTION, YEAR_OF_GRADUATION, SKILLS, UNIVERSITY, DEGREE) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING ID`
	err := db.QueryRow(query, user.Name, user.Email, user.Password, user.Description, user.YearOfGraduation, pq.StringArray(user.Skills), user.University, user.Degree).Scan(&id)

	return id, err
}

func selectAllFromStudentWhereEmailIs(db *sql.DB, email string) (dto.Student, error) {
	if db == nil {
		panic("db cannot be nil")
	}

	var user dto.Student
	query := `SELECT * FROM "STUDENT" WHERE EMAIL = $1`
	err := db.QueryRow(query, email).Scan(&user)
	if err != nil {
		return user, err
	}

	return user, err
}

func selectPasswordFromStudentWhereEmailIDs(db *sql.DB, email string) (string, error) {
	if db == nil {
		panic("db cannot be nil")
	}
	var password string
	query := `SELECT PASSWORD FROM "STUDENT" WHERE EMAIL = $1`
	err := db.QueryRow(query, email).Scan(&password)

	return password, err
}
