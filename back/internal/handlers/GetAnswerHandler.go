package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/models"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetAnswerHandler(c *gin.Context) {
	fmt.Println("\n[*] Received - GetAnswerHandler")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("[!] Error reading body! ", err.Error())
		log.Fatal(err)
	}

	fmt.Println("[*] Unmarshaling body")
	bodyUnmarshal := &models.GetAnswerRequest{}
	json.Unmarshal(body, &bodyUnmarshal)

	fmt.Println("[*] Checking if body is empty")
	if bodyUnmarshal.Message == "" {
		fmt.Println("[!] Body is empty!")
		c.IndentedJSON(http.StatusInternalServerError, "Message field is empty")
		return
	}

	fmt.Println("[*] Getting answer with algorithm")
	a := utils.GetBotAnswer(bodyUnmarshal.Message)

	c.IndentedJSON(http.StatusOK, a)
}
