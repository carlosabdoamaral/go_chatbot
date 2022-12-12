package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/models"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/rabbit"
	"github.com/gin-gonic/gin"
)

func SendKnowledgeListHandler(c *gin.Context) {
	fmt.Println("\n[*] Received SendKnowledgeListHandler")

	fmt.Println("[*] Reading list from body")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("[*] Error reading body! ", err.Error())
		log.Fatalln(err)
	}

	fmt.Println("[*] Unmarshalling body")
	bodyUnmarshal := []*models.Message{}
	json.Unmarshal(body, &bodyUnmarshal)

	fmt.Println("[*] Sending message to RabbitMQ server")
	rabbit.SendMessage(body)

	c.IndentedJSON(http.StatusOK, "SUCCESS")
}
