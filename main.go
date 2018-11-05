// This is a (school) project by me (https://niels-dingsbums.de) and Sheesher (https://ichbindumm12321.de)
package main

import (
	"github.com/gorilla/mux"
	"informatik_projekt/tictactoe"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tictactoe/{roomId}/getState/{cellId}", tictactoe.GetState).Methods("GET")
	r.HandleFunc("/tictactoe/{roomId}/updateState", tictactoe.UpdateState).Methods("POST")
	r.HandleFunc("/tictactoe/createRoom", tictactoe.CreateRoom).Methods("POST")

	go log.Fatal(http.ListenAndServe(":8080", r))
}
