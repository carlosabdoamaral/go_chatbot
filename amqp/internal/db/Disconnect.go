package db

import (
	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/global"
)

func Disconnect() {
	global.Database.Close()
}
