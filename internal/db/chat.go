package db

import (
	"database/sql"

	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
)

func SaveChatToDb(db *sql.DB, message dto.Message) error {
	return insertIntoChat(db, message)
}

func ViewChat(db *sql.DB, code string) ([]dto.Message, error) {
	return selectFromChatWhereCodeIs(db, code)
}

func insertIntoChat(db *sql.DB, message dto.Message) error {
	query := `INSERT INTO "CHAT" (SENDER_ID, CHAT_CODE, MESSAGE) VALUES ($1, $2, $3)`
	err := db.QueryRow(query, message.SenderId, message.ChatCode, message.Message).Err()
	return err
}

func selectFromChatWhereCodeIs(db *sql.DB, code string) ([]dto.Message, error) {
	query := `SELECT * FROM "CHAT" WHERE CHAT_CODE = $1`

	var messages []dto.Message

	rows, err := db.Query(query, code)
	if nil != err {
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		var message dto.Message

		if err := rows.Scan(
			&message.Id,
			&message.SenderId,
			&message.ChatCode,
			&message.Message,
			&message.SentAt,
		); nil != err {
			return []dto.Message{}, err
		}

		messages = append(messages, message)
	}

	return messages, err
}
