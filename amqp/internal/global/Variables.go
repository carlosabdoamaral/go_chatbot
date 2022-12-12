package global

import (
	"database/sql"

	"github.com/streadway/amqp"
)

var (
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
