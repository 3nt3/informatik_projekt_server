package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Global vars
var state GameState = GameState{[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}}

func GetState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(state.cells)
}

func UpdateState(w http.ResponseWriter, r *http.Request) {
	var data GameState
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(data)
}
