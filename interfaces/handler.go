package interfaces

import (
	"github.com/gorilla/mux"
	"main/interfaces/util"
	"net/http"
)

type HandlerPong struct {
}

func NewHandlerPong() HandlerPong {
	return HandlerPong{}
}

func (handler HandlerPong) GetPong(respuesta http.ResponseWriter, peticion *http.Request) {

	v := mux.Vars(peticion)

	util.JSON(respuesta,200,v)
}


