package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"math"
	"net/http"
)

// Global vars
var state GameState = GameState{[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}}

func GetState(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n=== GET  REQUEST ===")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(state.cells)
}

func UpdateState(w http.ResponseWriter, r *http.Request) {
	bar := pb.StartNew(9)

	var data [9]int
	_ = json.NewDecoder(r.Body).Decode(&data)
	state.cells = data
	fmt.Println("\n===  POST REQUEST  ===")

	for i, cell := range state.cells {
		if math.Mod(float64(i)+1.0, 3.0) == 0 {
			switch cell {
			case 0:
				fmt.Println("- ")
			case 1:
				fmt.Println("X ")
			case 2:
				fmt.Println("O ")
			}
		} else {
			switch cell {
			case 0:
				fmt.Print("- ")
			case 1:
				fmt.Print("X ")
			case 2:
				fmt.Print("O ")
			}
		}
		bar.Increment()
	}
	fmt.Println("")

}
