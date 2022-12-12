package common

import (
	"database/sql"

	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/models"
	"github.com/streadway/amqp"
)

var (
	IntentsFolder    = "ext/intents"
	Data             = []*models.Message{}
	History          = []*models.Message{}
	UnRecognizedWord = "Obrigado por me ensinar algo hoje 🥰"
	CharsToTrim      = []string{"!", "?", ".", "@"}
	Topics           = []string{}
	MockMode         = false
	Database         = &sql.DB{}
	IsRunningLocally = true
	ApiPort          = ""
)

var (
	DbUser         = ""
	DbPass         = ""
	DbHost         = ""
	DbName         = ""
	DbPort         = ""
	DbSsl          = ""
	DbUrl          = ""
	PostgresDriver = "postgres"
)

var (
	RabbitUrl        = ""
	RabbitQueueName  = ""
	RabbitConnection = &amqp.Connection{}
	RabbitChannel    = &amqp.Channel{}
	RabbitQueue      = &amqp.Queue{}
	RabbitPort       = ""
)
