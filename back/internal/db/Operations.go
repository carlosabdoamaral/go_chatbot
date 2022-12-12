package db

import (
	"fmt"

	global "github.com/Carlosabdoamaral/go_chatbot_backend/common"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/models"
	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/utils"
)

func ResetDatabase() {
	db := global.Database
	db.Exec("DELETE FROM message WHERE 1=1")
}

func SelectAllMessages() ([]*models.Message, error) {
	db := global.Database
	res := []*models.Message{}

	rows, err := db.Query(SelectAllMessagesQuery)
	if err != nil {

		return nil, err
	}

	for rows.Next() {
		m := &models.Message{}
		err = rows.Scan(
			&m.Message,
			&m.Answer,
		)

		m.Message = utils.DecodeString(m.Message)
		m.Answer = utils.DecodeString(m.Answer)

		if err != nil {
			return nil, err
		}

		res = append(res, m)
	}

	global.Data = res
	return res, nil
}

func SelectMessageById() (*models.Message, error) {
	db := global.Database
	res := &models.Message{}

	stmt := fmt.Sprintf("%s%s", SelectMessageByIdQuery, "1")
	row := db.QueryRow(stmt)
	fmt.Println(row)

	return res, nil
}
