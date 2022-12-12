package utils

import (
	"flag"

	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
)

func GetFlags() {
	isLocal := flag.Bool("local", true, "Run the API locally")
	flag.Parse()

	global.IsRunningLocally = *isLocal
}
