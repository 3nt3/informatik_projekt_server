// This is the "rock-paper-scissors" part of the API.
// It is used to handle rps requests
package rock_paper_scissors

import (
	"encoding/json"
	"log"
	"net/http"
)

// Global vars
var rooms []room
var players []player

// Create rps room
func CreateRoom(w http.ResponseWriter, r *http.Request) {

	var playerNames []string
	err := json.NewDecoder(r.Body).Decode(&playerNames)

	playersInRoom := []player{player{len(players), playerNames[0], 0}, player{len(players) + 1, playerNames[1], 0}}

	if err != nil {
		log.Println("Some json error just occurred! Call Bob the builder!!!")
	} else {
		room := room{len(rooms), playersInRoom}
		rooms = append(rooms, room)
		log.Printf("Added (rps) room: %v\n", room)
		json.NewEncoder(w).Encode(room.id)
	}

	//rooms = append()
}

/*
Post a figure (rock, paper or scissors)
The chosen figure is encoded like this:
0 - rock
1 - paper
2 - scissors
*/
func PostFigure(w http.ResponseWriter, r *http.Request) {

}

// Get score
func GetScore(w http.ResponseWriter, r *http.Request) {

}
