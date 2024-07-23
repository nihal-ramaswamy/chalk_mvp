package db

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
)

func AppendEmailToBookmarks(db *sql.DB, email string, addBookmark *dto.Bookmark) error {
	return updateBookmarksSetStudentEmailsArrayAppendWhereEmailIs(db, email, addBookmark)
}

func GetBookmarksForUser(db *sql.DB, email string) ([]string, error) {
	return selectStudentEmailsFromBookmarksWhereEmailIs(db, email)
}

// ----- Queries----

// TODO: Use student IDs instead of email
func updateBookmarksSetStudentEmailsArrayAppendWhereEmailIs(db *sql.DB, email string, addBookmark *dto.Bookmark) error {
	if nil == db {
		panic("db cannot be nil")
	}

	query := `SELECT COUNT(*) FROM "BOOKMARKS" WHERE EMAIL = $1`
	var count int
	err := db.QueryRow(query, email).Scan(&count)

	if nil != err {
		return err
	}

	if count == 0 {
		query = `INSERT INTO "BOOKMARKS" (EMAIL, STUDENT_EMAILS) VALUES ($1, ARRAY_APPEND(array[]::varchar[], $2))`
		err = db.QueryRow(query, email, addBookmark.StudentEmail).Err()
	} else {
		query = `UPDATE "BOOKMARKS" SET STUDENT_EMAILS = ARRAY_APPEND(STUDENT_EMAILS, $2) WHERE EMAIL = $1`
		err = db.QueryRow(query, email, addBookmark.StudentEmail).Err()
	}
	return err
}

func selectStudentEmailsFromBookmarksWhereEmailIs(db *sql.DB, email string) ([]string, error) {
	if nil == db {
		panic("db cannot be nil")
	}

	query := `SELECT STUDENT_EMAILS FROM "BOOKMARKS" WHERE EMAIL = $1`
	var bookmarks []string
	err := db.QueryRow(query, email).Scan(pq.Array(&bookmarks))

	return bookmarks, err
}
