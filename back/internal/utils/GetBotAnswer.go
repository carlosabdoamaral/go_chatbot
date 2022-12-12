package utils

import (
	"math/rand"
	"strings"

	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/models"
)

func GetBotAnswer(s string) *models.Message {
	var (
		data            = global.Data
		possibleAnswers = []*models.Message{}
		res             = &models.Message{}
	)

	for _, message := range data {
		if strings.EqualFold(s, message.Message) {
			possibleAnswers = append(possibleAnswers, message)
		}
	}

	if len(possibleAnswers) == 0 {
		return &models.Message{
			Message: s,
			Answer:  global.UnRecognizedWord,
		}
	} else {
		k := rand.Intn(len(possibleAnswers))
		res = possibleAnswers[k]

		return &models.Message{
			Message: strings.ToUpper(s),
			Answer:  strings.ToUpper(CapitalizeWord(res.Answer)),
		}
	}
}
