package db

import (
	"strings"

	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/global"
	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/models"
)

func SendKnowledge(k *models.Message) error {
	db := global.Database

	invalidChars := []string{"?", "!", "@", "#", "$", "%", "ˆ", "&", "*", "(", ")", "."}

	for _, letter := range k.Message {
		currentLetter := string(letter)

		for _, c := range invalidChars {
			if currentLetter == c {
				k.Message = strings.Replace(k.Message, currentLetter, "", 1)
			}
		}
	}

	k.Message = strings.Replace(k.Message, "?", "", -1)
	k.Message = strings.Replace(k.Message, "!", "", -1)
	k.Message = strings.Replace(k.Message, ".", "", -1)

	insert, err := db.Prepare(InsertMessageQuery)
	checkErr(err)

	_, err = insert.Exec(k.Message, k.Answer)
	checkErr(err)

	return nil
}
