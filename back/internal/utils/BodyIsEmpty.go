package utils

import (
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/models"
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
