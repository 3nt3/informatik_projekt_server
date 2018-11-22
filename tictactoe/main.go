// Handle tic tac toe requests
package tictactoe

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Global vars
var rooms []room
var players []player

func GetState(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("\n=== GET  REQUEST ===")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	roomId, _ := strconv.Atoi(mux.Vars(r)["roomId"])

	room := rooms[roomId]
	cells := room.state.cells

	log.Printf("(ttt) Get state (room %d): %d\n", room.id, cells)

	_ = json.NewEncoder(w).Encode(cells)
}

func UpdateState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	roomId, _ := strconv.Atoi(mux.Vars(r)["roomId"])

	var state []int
	_ = json.NewDecoder(r.Body).Decode(&state)

	rooms[roomId].state = gameState{state}
	log.Printf("(ttt) New State for room %d %v\n", roomId, state)
}

// Create room
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var data []string
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		log.Println("(ttt) Something went wrong with json decoding!!! Call the Hydrauliknotdienst!")
	} else {
		playersInRoom := []player{{len(players), data[0], 0}, {len(players) + 1, data[1], 0}}

		room := room{len(rooms), playersInRoom, gameState{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0}}}
		rooms = append(rooms, room)
		//fmt.Println(rooms)
		for _, player := range playersInRoom {
			players = append(players, player)
		}

		log.Printf("(ttt) New room: %v\n", room)

		json.NewEncoder(w).Encode(room.id)
	}
}

func GetScores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	roomId, _ := strconv.Atoi(mux.Vars(r)["roomId"])

	playersInRoom := rooms[roomId].players
	var scores []int

	for _, player := range playersInRoom {
		scores = append(scores, player.score)
	}

	json.NewEncoder(w).Encode(scores)

}

func UpdateScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	roomId, _ := strconv.Atoi(mux.Vars(r)["roomId"])
	playerId, _ := strconv.Atoi(mux.Vars(r)["playerId"])

	currentPlayer := rooms[roomId].players[playerId]

	var newScore int
	_ = json.NewDecoder(r.Body).Decode(&newScore)

	currentPlayer.score = newScore

}

// Generate a random state
func Random(w http.ResponseWriter, r *http.Request) {
	roomId, _ := strconv.Atoi(mux.Vars(r)["roomId"])

	playersInRoom := rooms[roomId].players

	var state []int

	for i := 0; i < 9; i++ {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		x := r.Intn(3)
		state = append(state, x)
	}
	// fmt.Println(state)

	log.Printf("(ttt) It's time for some more randomness in room %d %v\n", roomId, playersInRoom)

}

func Keksen(w http.ResponseWriter, r *http.Request) {
	roomId, _ := strconv.Atoi(mux.Vars(r)["roomId"])

	//playersInRoom := rooms[roomId].players
	var scores []int

	for i := 0; i < 2; i++ {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		x := r.Intn(10000)
		scores = append(scores, x)
	}
	log.Printf("(ttt) Keeeeeeeeeeeksen in room %d (%d, %d)", roomId, scores[0], scores[1])

	rooms[roomId].players[0].score = scores[0]
	rooms[roomId].players[1].score = scores[1]

}
