package db

import (
	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/global"
)

func InitDatabase() error {
	db := global.Database
	db.Exec("INSERT INTO message(message, answer) VALUES ('OI', 'EAI, TUDO BEM?');")

	return nil
}
