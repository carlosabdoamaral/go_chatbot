package db

import "fmt"

var (
	SelectAllMessagesQuery = "SELECT message, answer FROM message"
	SelectMessageByIdQuery = "SELECT message, answer FROM message WHERE id="
	InsertMessageQuery     = fmt.Sprintf("INSERT INTO message(message, answer) VALUES ($1,$2)")
)
