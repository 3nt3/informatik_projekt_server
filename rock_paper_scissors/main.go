// Handle rock paper scissors requests
package rock_paper_scissors

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Global vars
var rooms []room
var players []player

// Create rps room
func CreateRoom(w http.ResponseWriter, r *http.Request) {

	var playerNames []string
	err := json.NewDecoder(r.Body).Decode(&playerNames)

	playersInRoom := []player{player{len(players), playerNames[0], 0, -1}, player{len(players) + 1, playerNames[1], 0, -1}}

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
	var fig int
	json.NewDecoder(r.Body).Decode(&fig)

	playerId, _ := strconv.Atoi(mux.Vars(r)["playerId"])
	roomId, _ := strconv.Atoi(mux.Vars(r)["roomId"])

	rooms[roomId].players[playerId].figure = fig

	log.Printf("Updated fig of player \"%s\" to %d\n", rooms[roomId].players[playerId].name, rooms[roomId].players[playerId].figure)
}

func GetFigures(w http.ResponseWriter, r *http.Request) {
	roomId, _ := strconv.Atoi(mux.Vars(r)["roomId"])

	var figs []int

	for _, player := range rooms[roomId].players {
		figs = append(figs, player.figure)
	}

	json.NewEncoder(w).Encode(figs)
	log.Printf("GET figs of room %d %v -> %d\n", rooms[roomId].id, rooms[roomId].players, figs)
}

// Get score
func GetScores(w http.ResponseWriter, r *http.Request) {
	roomId, _ := strconv.Atoi(mux.Vars(r)["roomId"])

	var scores []int
	for _, player := range rooms[roomId].players {
		scores = append(scores, player.score)
	}

	json.NewEncoder(w).Encode(scores)

	log.Println("")
}
