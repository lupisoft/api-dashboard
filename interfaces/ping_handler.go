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

func (handler HandlerPong) GetPong(response http.ResponseWriter, request *http.Request) {
	util.Respond(response, 200, "pong")
}
