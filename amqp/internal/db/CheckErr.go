package db

import "log"

func checkErr(err error) error {
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
