package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/models"
)

func InitConversation() {
	var (
		possibleAnswers = []*models.Message{}
		answer          = ""
	)

	// for _, m := range global.Data {
	// 	if m.Category == "GREETINGS" {
	// 		possibleAnswers = append(possibleAnswers, m)
	// 	}
	// }

	randIndex := rand.Intn(len(possibleAnswers))
	answer = possibleAnswers[randIndex].Answer
	answer = CapitalizeWord(answer)
	fmt.Printf("[IA  ] -> %s\n", answer)

	global.History = append(global.History, possibleAnswers[randIndex])
}

func Conversation() {
	var (
		userInput = ""
		scanner   = bufio.NewScanner(os.Stdin)
	)

	for {
		fmt.Print("[YOU ] -> ")
		scanner.Scan()
		userInput = scanner.Text()

		for _, mark := range global.CharsToTrim {
			userInput = strings.Trim(userInput, mark)
		}

		answer := GetBotAnswer(userInput)
		global.History = append(global.History, answer)

		fmt.Printf("[IA  ] -> %s\n", CapitalizeWord(answer.Answer))
	}
}
