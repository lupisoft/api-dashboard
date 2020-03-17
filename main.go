package main

import (
	"fmt"
	_ "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"log"
	"main/interfaces"
	"net/http"
	"time"
)

func main()  {

	enrutador := mux.NewRouter()
	direccion := ":8000"

	servidor := &http.Server{
		Handler: enrutador,
		Addr:    direccion,
		// Timeouts para evitar que el servidor se quede "colgado" por siempre
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	handler := interfaces.NewHandlerPong()
	enrutador.HandleFunc("/ping", handler.GetPong ).Methods("GET")

	fmt.Printf("Escuchando en %s. Presiona CTRL + C para salir", direccion)
	log.Fatal(servidor.ListenAndServe())


}




