package main

import (
	"fmt"

	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/db"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/handlers"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/middlewares"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/rabbit"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/utils"
	"github.com/gin-gonic/gin"
)

// COMMAND : go run main.go --local true
func main() {
	fmt.Println("[*] Starting application")

	InitVariables()
	fmt.Println("[*] Variables are initialized!")

	InitRabbit()
	fmt.Println("[*] Rabbit Connection is open!")

	InitDatabase()
	fmt.Println("[*] Database Connection is open!")

	RunApi()
}

func InitVariables() {
	utils.GetFlags()
	utils.GetEnvVariables()
	utils.HandleIsRunningLocally()
}

func InitRabbit() {
	rabbit.Connect()
	rabbit.ConnectToChannel()
}

func InitDatabase() {
	db.Connect()
	db.InitDatabase()
	db.ResetDatabase()
	db.SelectAllMessages()
}

func RunApi() {
	router := gin.Default()
	router.Use(middlewares.CORS())

	routerWithPrefix := router.Group("/api")
	routerWithPrefix.POST("/conversation", handlers.GetAnswerHandler)

	routerWithPrefix.POST("/knowledge/new", handlers.SendKnowledgeHandler)
	routerWithPrefix.POST("/knowledge/new/list", handlers.SendKnowledgeListHandler)
	routerWithPrefix.GET("/knowledge/all", handlers.GetKnowledgeHandler)

	router.Run(fmt.Sprintf(":%s", global.ApiPort))
}
