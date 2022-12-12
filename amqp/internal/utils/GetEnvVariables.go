package utils

import (
	"log"
	"os"

	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/global"
	"github.com/joho/godotenv"
)

func GetEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to load env file. Make sure .env file exists!")
	}

	global.RabbitPort = os.Getenv("RABBIT_PORT")
	global.RabbitQueueName = os.Getenv("RABBIT_QUEUE_NAME")
	global.RabbitUrl = os.Getenv("RABBIT_URL")
	global.ApiPort = os.Getenv("API_PORT")
	global.DbUser = os.Getenv("DB_USER")
	global.DbPass = os.Getenv("DB_PASSWORD")
	global.DbHost = os.Getenv("DB_HOST")
	global.DbName = os.Getenv("DB_NAME")
	global.DbPort = os.Getenv("DB_PORT")
	global.DbSsl = os.Getenv("DB_SSL")
	global.RabbitUrl = os.Getenv("RABBIT_URL")
	global.DbUrl = ""
}
