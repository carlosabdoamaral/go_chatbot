package rabbit

import (
	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
	"github.com/streadway/amqp"
)

func SendMessage(msg []byte) error {
	err := global.RabbitChannel.Publish(
		"",
		"knowledgeQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
