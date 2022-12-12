package rabbit

import (
	"log"

	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
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
