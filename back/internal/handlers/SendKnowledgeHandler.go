package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/models"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/rabbit"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func SendKnowledgeHandler(c *gin.Context) {
	fmt.Println("\n[*] Received - SendKnowledgeHandler")

	fmt.Println("[*] Reading body")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("[!] Error reading body! ", err.Error())
		log.Fatal(err)
	}

	fmt.Println("[*] Unmarshaling body")
	bodyUnmarshal := &models.Message{}
	json.Unmarshal(body, &bodyUnmarshal)

	fmt.Println("[*] Checking if body is empty")
	bodyIsEmpty := utils.BodyIsEmpty(bodyUnmarshal)

	if bodyIsEmpty {
		fmt.Println("[!] Body is empty!")
		c.IndentedJSON(http.StatusInternalServerError, "One or more fields are empty")
		return
	}
	err = rabbit.SendMessage(body)
	if err != nil {
		fmt.Println("[!] Rabbit error! ", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "Success")

}
