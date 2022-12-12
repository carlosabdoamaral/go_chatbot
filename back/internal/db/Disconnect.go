package db

import (
	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
)

func Disconnect() {
	global.Database.Close()
}
