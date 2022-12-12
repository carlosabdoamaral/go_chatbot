package utils

import (
	"fmt"
	"os"

	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
)

func HandleIsRunningLocally() {
	if global.IsRunningLocally {
		global.ApiPort = "8080"
		global.DbSsl = "disable"
		global.DbUrl = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", global.DbHost, global.DbPort, global.DbUser, global.DbPass, global.DbName, global.DbSsl)
	} else {
		global.DbUrl = os.Getenv("DATABASE_URL")
	}
}
