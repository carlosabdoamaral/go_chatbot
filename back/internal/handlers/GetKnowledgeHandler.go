package handlers

import (
	"fmt"
	"net/http"

	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/db"
	"github.com/gin-gonic/gin"
)

func GetKnowledgeHandler(c *gin.Context) {
	fmt.Println("\n[*] Received - GetKnowledgeHandler")
	fmt.Println("[*] Getting messages from database")

	_, err := db.SelectAllMessages()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusOK, global.Data)
}
