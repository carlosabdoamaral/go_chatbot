package utils

import (
	"log"
	"os"

	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
	"github.com/joho/godotenv"
)

func GetEnvVariables() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Println("Failed to load env file. Make sure .env file exists!")
	}

	global.ApiPort = os.Getenv("PORT")
	global.DbUser = os.Getenv("DB_USER")
	global.DbPass = os.Getenv("DB_PASSWORD")
	global.DbHost = os.Getenv("DB_HOST")
	global.DbName = os.Getenv("DB_NAME")
	global.DbPort = os.Getenv("DB_PORT")
	global.DbSsl = os.Getenv("DB_SSL")
	global.RabbitUrl = os.Getenv("RABBIT_URL")
	global.DbUrl = ""
}
