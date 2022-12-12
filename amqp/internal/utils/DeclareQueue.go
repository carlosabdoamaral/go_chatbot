package utils

import (
	"log"

	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/global"
)

func DeclareQueue() {
	ch, err := global.RabbitConnection.Channel()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	q, err := ch.QueueDeclare(
		global.RabbitQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	global.RabbitChannel = ch
	global.RabbitQueue = &q
}
