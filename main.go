// This is a (school) project by me (https://niels-dingsbums.de) and Sheesher (https://ichbindumm12321.de)
package main

import (
	"informatik_projekt_server/rock_paper_scissors"
	"informatik_projekt_server/tictactoe"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()	
	// tic-tac-toe section
	r.HandleFunc("/tictactoe/{roomId}/getState", tictactoe.GetState).Methods("GET")
	r.HandleFunc("/tictactoe/{roomId}/updateState", tictactoe.UpdateState).Methods("POST")
	r.HandleFunc("/tictactoe/createRoom", tictactoe.CreateRoom).Methods("POST")	
	r.HandleFunc("/tictactoe/{roomId}/getScores", tictactoe.GetScores).Methods("GET")
	r.HandleFunc("/tictactoe/{roomId/updateScore/{playerId}", tictactoe.UpdateScore).Methods("POST")

	// rock-paper-scissors section
	r.HandleFunc("/rps/{roomId}/getScores", rock_paper_scissors.GetScores).Methods("GET")
	r.HandleFunc("/rps/{roomId}/updateScore/{playerId}", rock_paper_scissors.UpdateScore).Methods("POST")
	r.HandleFunc("/rps/{roomId}/postFig/{playerId}", rock_paper_scissors.PostFigure).Methods("POST")
	r.HandleFunc("/rps/{roomId}/getFigs", rock_paper_scissors.GetFigures).Methods("GET")
	r.HandleFunc("/rps/createRoom", rock_paper_scissors.CreateRoom).Methods("POST")

	go log.Fatal(http.ListenAndServe(":8000", r))
}
