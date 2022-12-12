package main

import (
	"fmt"

	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/db"
	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/global"
	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/handlers"
	"github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/utils"
)

func main() {
	InitVariables()
	InitDatabase()
	InitRabbit()

	defer global.Database.Close()
	defer global.RabbitConnection.Close()
	defer global.RabbitChannel.Close()
}

func InitVariables() {
	utils.GetFlags()
	utils.GetEnvVariables()
	utils.HandleIsRunningLocally()
}

func InitRabbit() {
	utils.Connect()
	fmt.Println("[*] Rabbit Connection is open!")

	utils.DeclareQueue()
	fmt.Println("[*] Rabbit Queue is declared!")

	fmt.Println("[*] Rabbit Connected!")
	handlers.Consumer()
}

func InitDatabase() {
	db.Connect()
	fmt.Println("[*] Database Connected!")
}

// func _test() {
// 	s := "EAI MUNDO"
// 	fmt.Println(string(s))

// 	s = utils.EncodeString(s)
// 	fmt.Println(string(s))

// 	s = utils.DecodeString(s)
// 	fmt.Println(string(s))
// }
