package rock_paper_scissors

// Rooms are used as areas for two players to play in
type room struct {
	id      int
	players []player
}

// The player type
type player struct {
	id     int
	name   string
	score  int
	figure int
}
