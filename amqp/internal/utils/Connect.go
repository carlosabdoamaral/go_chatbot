package utils

import (
	"log"

	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/global"
	"github.com/streadway/amqp"
)

func Connect() {
	conn, err := amqp.Dial(global.RabbitUrl)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	global.RabbitConnection = conn
}
