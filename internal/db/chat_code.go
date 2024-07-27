package db

import (
	"database/sql"
	"fmt"

	"github.com/nihal-ramaswamy/chalk_mvp/internal/utils"
)

func CreateNewChatCode(db *sql.DB, id1, id2, code string) error {
	return insertIntoChatCode(db, id1, id2, code)
}

func DoesCodeExist(db *sql.DB, id1, id2 string) (bool, error) {
	return selectCountFromChatCodeWhereId1IsAndId2Is(db, id1, id2)
}

func GetCode(db *sql.DB, id1, id2 string) (string, error) {
	return selectCodeFromChatCodeWhereId1IsAndId2Is(db, id1, id2)
}

func IsStudentAuthorizedToViewChat(db *sql.DB, id, code string) (bool, error) {
	ok, err := selectCountFromChatCodeWhereExistsSelectCodeFromChatCodeWhereCodeIsAndId1Is(db, id, code)
	if nil != err {
		return false, err
	}

	if ok {
		return ok, nil
	}

	ok, err = selectCountFromChatCodeWhereExistsSelectCodeFromChatCodeWhereCodeIsAndId2Is(db, id, code)
	return ok, err
}

// ---- Queries ----
func insertIntoChatCode(db *sql.DB, id1, id2, code string) error {
	if id1 < id2 {
		utils.Swap(&id1, &id2)
	}
	query := `INSERT INTO "CHAT_CODE" (ID1, ID2, CODE) VALUES ($1, $2, $3)`
	err := db.QueryRow(query, id1, id2, code).Err()
	return err
}

func selectCountFromChatCodeWhereId1IsAndId2Is(db *sql.DB, id1, id2 string) (bool, error) {
	if id1 < id2 {
		utils.Swap(&id1, &id2)
	}

	query := `SELECT COUNT(*) FROM "CHAT_CODE" WHERE ID1 = $1 and ID2 = $2`
	var count int
	err := db.QueryRow(query, id1, id2).Scan(&count)

	if nil != err {
		return false, err
	}

	exists := (count > 0)

	return exists, err
}

func selectCodeFromChatCodeWhereId1IsAndId2Is(db *sql.DB, id1, id2 string) (string, error) {
	if id1 < id2 {
		utils.Swap(&id1, &id2)
	}

	query := `SELECT COUNT(*) FROM "CHAT_CODE" WHERE ID1 = $1 and ID2 = $2`
	var count int
	err := db.QueryRow(query, id1, id2).Scan(&count)

	if nil != err {
		return "", err
	}

	if count > 0 {
		query := `SELECT CODE FROM "CHAT_CODE" WHERE ID1 = $1 and ID2 = $2`
		var code string
		err := db.QueryRow(query, id1, id2).Scan(&code)

		return code, err
	}

	return "", fmt.Errorf("error fetching code")
}

func selectCountFromChatCodeWhereExistsSelectCodeFromChatCodeWhereCodeIsAndId1Is(db *sql.DB, id, code string) (bool, error) {
	query := `SELECT COUNT(*) FROM "CHAT_CODE" WHERE CODE = $1 AND ID1 = $2`
	var count int
	err := db.QueryRow(query, code, id).Scan(&count)
	return count > 0, err
}

func selectCountFromChatCodeWhereExistsSelectCodeFromChatCodeWhereCodeIsAndId2Is(db *sql.DB, id, code string) (bool, error) {
	query := `SELECT COUNT(*) FROM "CHAT_CODE" WHERE CODE = $1 AND ID2 = $2`
	var count int
	err := db.QueryRow(query, code, id).Scan(&count)
	return count > 0, err
}
