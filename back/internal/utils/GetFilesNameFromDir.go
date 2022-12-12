package utils

import (
	"fmt"
	"io/ioutil"
	"log"

	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
)

func GetFilesNameFromDir() []string {
	filesList := []string{}

	files, err := ioutil.ReadDir(global.IntentsFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			nameWithPrefix := fmt.Sprintf("%s/%s", global.IntentsFolder, file.Name())
			filesList = append(filesList, nameWithPrefix)
		}
	}
	return filesList
}
