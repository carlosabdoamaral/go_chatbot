package utils

import (
	"flag"

	global "github.com/Carlosabdoamaral/go_chatbot_rabbit/internal/global"
)

func GetFlags() {
	isLocal := flag.Bool("local", true, "Run the API locally")
	flag.Parse()

	global.IsRunningLocally = *isLocal
}
