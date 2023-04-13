package handler

import (
	"errors"
	"log"
)

func PanicIfError(err error) {
	if err != nil {
		err = errors.New("error--" + err.Error())
		log.Println(err)
		panic(err)
	}
}
