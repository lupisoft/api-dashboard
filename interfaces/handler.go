package interfaces

import (
	"main/interfaces/util"
	"net/http"
)

type HandlerPong struct {
}

func NewHandlerPong() HandlerPong {
	return HandlerPong{}
}

func (handler HandlerPong) GetPong(respuesta http.ResponseWriter, peticion *http.Request) {

	result := "pong"

	util.JSON(respuesta,200,result)
}


