package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tictactoe_server/api"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tictactoe/getState", api.GetState).Methods("GET")
	r.HandleFunc("/tictactoe/updateState", api.UpdateState).Methods("POST")

	go log.Fatal(http.ListenAndServeTLS())
}
