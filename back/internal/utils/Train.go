package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/models"
)

func TrainData(files []string) []*models.Message {
	res := []*models.Message{}
	for _, file := range files {
		log.Printf("Trainning %s", file)

		jsonFile, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
		}

		var messages []*models.Message
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &messages)

		for _, m := range messages {
			obj := &models.Message{
				Message: strings.ToUpper(m.Message),
				Answer:  strings.ToUpper(m.Answer),
			}

			res = append(res, obj)
		}

		defer jsonFile.Close()
	}

	return res
}
