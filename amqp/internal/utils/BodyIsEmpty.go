package utils

import (
	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/models"
)

func BodyIsEmpty(m *models.Message) bool {
	messageIsEmpty := m.Message == ""
	answerIsEmpty := m.Answer == ""

	if messageIsEmpty || answerIsEmpty {
		return true
	} else {
		return false
	}
}
