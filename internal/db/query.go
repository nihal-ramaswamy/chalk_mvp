package db

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
)

func insertIntoConference(db *sql.DB, conference *dto.Conference) error {
	if nil == db {
		return errors.New("sql db nil")
	}

	query := `INSERT INTO "CONFERENCE" (CODE, ADMIN, ACTIVE) VALUES ($1, $2, $3)`
	err := db.QueryRow(query, conference.Code, conference.Admin, conference.Active).Err()

	return err
}

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

func updateBookmarksSetStudentEmailsArrayAppendWhereEmailIs(db *sql.DB, email string, addBookmark *dto.Bookmark) error {
	if nil == db {
		panic("db cannot be nil")
	}

	query := `UPDATE BOOKMARKS SET STUDENT_EMAILS = ARRAY_APPEND(STUDENT_EMAILS, $2) WHERE EMAIL = $1`
	err := db.QueryRow(query, email, addBookmark.StudentEmail).Err()
	return err
}

func selectStudentEmailsFromBookmarksWhereEmailIs(db *sql.DB, email string) ([]string, error) {
	if nil == db {
		panic("db cannot be nil")
	}

	query := `SELECT STUDENT_EMAILS FROM BOOKMARKS WHERE EMAIL = $1`
	var bookmarks []string
	err := db.QueryRow(query, email).Scan(&bookmarks)

	return bookmarks, err
}
