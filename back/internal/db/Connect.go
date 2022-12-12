package db

import (
	"database/sql"

	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open(global.PostgresDriver, global.DbUrl)
	if err != nil {
		return nil, err
	}

	global.Database = db
	return global.Database, nil
}
