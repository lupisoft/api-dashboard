package main

import (
	"github.com/gorilla/mux"
	"log"
	"main/config"
	"main/interfaces"
	"net/http"
	"time"
)

func main() {

	config.NewDBConnection()

	router := mux.NewRouter()
	port := ":8000"
	server := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	handler := interfaces.NewHandlerPong()
	router.HandleFunc("/ping", handler.GetPong).Methods("GET")

	log.Fatal(server.ListenAndServe())
}
