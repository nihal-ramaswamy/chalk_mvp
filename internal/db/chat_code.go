package db

import "database/sql"

func CreateNewChatCode(db *sql.DB, id1, id2, code string) error {
	return insertIntoChatCode(db, id1, id2, code)
}

func DoesCodeExist(db *sql.DB, id1, id2 string) (bool, error) {
	return selectCountFromChatCodeWhereId1IsAndId2Is(db, id1, id2)
}

// ---- Queries ----
func insertIntoChatCode(db *sql.DB, id1, id2, code string) error {
	query := `INSERT INTO "CHAT_CODE" (ID1, ID2, CODE) VALUES ($1, $2, $3)`
	err := db.QueryRow(query, id1, id2, code).Err()

	return err
}

func selectCountFromChatCodeWhereId1IsAndId2Is(db *sql.DB, id1, id2 string) (bool, error) {
	query := `SELECT COUNT(*) FROM "CHAT_CODE" WHERE ID1 = $1 and ID2 = $2`
	var count1 int
	err := db.QueryRow(query, id1, id2).Scan(&count1)

	if nil != err {
		return false, err
	}

	var count2 int
	err = db.QueryRow(query, id2, id1).Scan(&count2)

	exists := (count1 > 0) || (count2 > 0)

	return exists, err
}
