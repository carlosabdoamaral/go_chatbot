package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/db"
	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/global"
	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/models"
	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/utils"
)

func Consumer() {
	msgs, err := global.RabbitChannel.Consume(
		global.RabbitQueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Error")
	}

	forever := make(chan bool)
	go func() {
		for m := range msgs {
			message, err := BuildMessageModel(m.Body)
			if err != nil {
				return
			}

			fmt.Println("\n[*] Received new message")
			fmt.Printf("[*] Message: %s - Answer: %s\n", message.Message, message.Answer)

			err = db.SendKnowledge(message)
			if err != nil {
				return
			}

			fmt.Println("[*] Message was sent to database successfully")
		}
	}()

	fmt.Println("[*] Waiting for messages")
	<-forever
}

func BuildMessageModel(body []byte) (*models.Message, error) {
	bodyUnmarshal := &models.Message{}
	json.Unmarshal(body, &bodyUnmarshal)

	bodyIsEmpty := utils.BodyIsEmpty(bodyUnmarshal)
	if bodyIsEmpty {
		return nil, errors.New("the body is empty")
	}

	res := &models.Message{
		Message: utils.EncodeString(strings.ToUpper(bodyUnmarshal.Message)),
		Answer:  utils.EncodeString(strings.ToUpper(bodyUnmarshal.Answer)),
	}

	return res, nil
}
