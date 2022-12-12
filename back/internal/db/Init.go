package db

import (
	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
)

func InitDatabase() error {
	db := global.Database
	// db.Exec("DROP DATABASE ia_using_go")
	// db.Exec("create database " + dbName)
	// db.Exec("DROP TABLE message;")
	db.Exec("CREATE TABLE message (id serial PRIMARY KEY,message VARCHAR(255) NOT NULL,answer VARCHAR(255) NOT NULL);")
	db.Exec("INSERT INTO message(message, answer) VALUES ('OI', 'EAI, TUDO BEM?');")

	return nil
}
