package tictactoe

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Global vars
var rooms []room
var players []player

func GetState(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("\n=== GET  REQUEST ===")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "https://snap.berkeley.edu")

	roomId, _ := strconv.Atoi(mux.Vars(r)["roomId"])
	cellId, _ := strconv.Atoi(mux.Vars(r)["cellId"])

	room := rooms[roomId]
	cell := room.state.cells[cellId]

	fmt.Println(rooms)
	fmt.Printf("Get state (room %d, cell %d): %d\n", room.id, cellId, cell)

	_ = json.NewEncoder(w).Encode(cell)
}

func UpdateState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "https://snap.berkeley.edu")
	roomId, _ := strconv.Atoi(mux.Vars(r)["roomId"])

	var state []int
	_ = json.NewDecoder(r.Body).Decode(&state)

	rooms[roomId].state = gameState{state}
	fmt.Printf("==> Actual state: %d\n\n", rooms[roomId].state)
}

// Create room
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "https://snap.berkeley.edu")
	var data []string
	_ = json.NewDecoder(r.Body).Decode(&data)

	playersInRoom := []player{player{len(players), data[0]}, player{len(players) + 1, data[1]}}

	room := room{len(rooms), playersInRoom, gameState{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0}}}
	rooms = append(rooms, room)
	fmt.Println(rooms)
	for _, player := range playersInRoom {
		players = append(players, player)
	}

	fmt.Printf("New room: %v\n", room)

	json.NewEncoder(w).Encode(room.id)
}
