package rabbit

import (
	"log"

	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
)

func ConnectToChannel() {
	ch, err := global.RabbitConnection.Channel()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	global.RabbitChannel = ch
}
