package models

type Message struct {
	Message string `json:"message"`
	Answer  string `json:"answer"`
}

type GetAnswerRequest struct {
	Message string `json:"message"`
}
