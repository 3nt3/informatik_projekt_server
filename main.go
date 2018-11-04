package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tictactoe_server/api"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tictactoe/{roomId}/getState/{cellId}", api.GetState).Methods("GET")
	r.HandleFunc("/tictactoe/{roomId}/updateState", api.UpdateState).Methods("POST")
	r.HandleFunc("/tictactoe/createRoom", api.CreateRoom).Methods("POST")

	go log.Fatal(http.ListenAndServe(":8080", r))
}
