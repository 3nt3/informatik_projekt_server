package tictactoe

type gameState struct {
	cells []int `json:"cells"`
}

type player struct {
	id   int
	name string
	score int
}

type room struct {
	id      int
	players []player
	state   gameState
}
