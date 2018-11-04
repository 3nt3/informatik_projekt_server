package api

type GameState struct {
	cells []int `json:"cells"`
}

type Player struct {
	id   int
	name string
}

type Room struct {
	id      int
	players []Player
	state   GameState
}
