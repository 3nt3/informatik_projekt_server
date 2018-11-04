package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Global vars
var Rooms []Room
var Players []Player

func GetState(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("\n=== GET  REQUEST ===")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "https://snap.berkeley.edu")

	roomId, _ := strconv.Atoi(mux.Vars(r)["roomId"])
	cellId, _ := strconv.Atoi(mux.Vars(r)["cellId"])

	room := Rooms[roomId]
	cell := room.state.cells[cellId]

	fmt.Println(Rooms)
	fmt.Printf("Get state (room %d, cell %d): %d\n", room.id, cellId, cell)

	_ = json.NewEncoder(w).Encode(cell)
}

func UpdateState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "https://snap.berkeley.edu")
	roomId, _ := strconv.Atoi(mux.Vars(r)["roomId"])

	var state []int
	_ = json.NewDecoder(r.Body).Decode(&state)

	Rooms[roomId].state = GameState{state}
	fmt.Printf("==> Actual state: %d\n\n", Rooms[roomId].state)
}

// Create room
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "https://snap.berkeley.edu")
	var data []string
	_ = json.NewDecoder(r.Body).Decode(&data)

	PlayersInRoom := []Player{Player{len(Players), data[0]}, Player{len(Players) + 1, data[1]}}

	room := Room{len(Rooms), PlayersInRoom, GameState{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0}}}
	Rooms = append(Rooms, room)
	fmt.Println(Rooms)
	for _, player := range PlayersInRoom {
		Players = append(Players, player)
	}

	fmt.Printf("New room: %v\n", room)

	json.NewEncoder(w).Encode(room.id)
}
