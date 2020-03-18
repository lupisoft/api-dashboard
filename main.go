package main

import (
	"github.com/gorilla/mux"
	"log"
	"main/interfaces"
	"net/http"
	"time"
)

func main() {
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
