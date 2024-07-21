package db

import (
	"database/sql"
	"errors"

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

func insertIntoUser(db *sql.DB, user *dto.Student) (string, error) {
	if db == nil {
		panic("db cannot be nil")
	}

	var id string
	query := `INSERT INTO "USER" (NAME, EMAIL, PASSWORD, DESCRIPTION, YEAR_OF_GRADUATION, SKILLS, UNIVERSITY) VALUES ($1, $2, $3) RETURNING ID`
	err := db.QueryRow(query, user.Name, user.Email, user.Password, user.Description, user.YearOfGraduation, user.Skills, user.University).Scan(&id)

	return id, err
}

func selectAllFromUserWhereEmailIs(db *sql.DB, email string) (dto.Student, error) {
	if db == nil {
		panic("db cannot be nil")
	}

	var user dto.Student
	query := `SELECT * FROM "USER" WHERE EMAIL = $1`
	err := db.QueryRow(query, email).Scan(&user)
	if err != nil {
		return user, err
	}

	return user, err
}

func selectPasswordFromUserWhereEmailIDs(db *sql.DB, email string) (string, error) {
	if db == nil {
		panic("db cannot be nil")
	}
	var password string
	query := `SELECT PASSWORD FROM "USER" WHERE EMAIL = $1`
	err := db.QueryRow(query, email).Scan(&password)

	return password, err
}
