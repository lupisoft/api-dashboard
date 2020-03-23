package main

import (
	"github.com/gorilla/mux"
	"log"
	"main/config"
	"main/interfaces"
	"net/http"
)

func main() {

	configuration := config.BuildConfiguration()

	router := mux.NewRouter()
	server := &http.Server{
		Handler:      router,
		Addr:         configuration.Server.Port,
		WriteTimeout: configuration.Server.WriteTimeOut,
		ReadTimeout:  configuration.Server.ReadTimeOut,
	}

	handler := interfaces.NewHandlerPong()
	router.HandleFunc("/ping", handler.GetPong).Methods("GET")

	log.Fatal(server.ListenAndServe())
}
