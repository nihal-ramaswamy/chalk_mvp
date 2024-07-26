package db

import (
	"database/sql"

	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
)

func SaveChatToDb(db *sql.DB, message dto.Message) error {
	return insertIntoChat(db, message)
}

func insertIntoChat(db *sql.DB, message dto.Message) error {
	query := `INSERT INTO "CHAT" (SENDER_ID, CHAT_CODE, MESSAGE) VALUES ($1, $2, $3)`
	err := db.QueryRow(query, message.SenderId, message.ChatCode, message.Message).Err()
	return err
}
