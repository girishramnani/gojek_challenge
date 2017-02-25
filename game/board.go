package game

// Board keeps track of the attacks which were done on it.
type Board struct {
	gridSize int
	ships    []Ship
	attacks  []Coord
}

func NewBoard(gridSize int, ships ...Ship) *Board {

	return &Board{
		gridSize: gridSize,
		ships:    ships,
		attacks:  make([]Coord, 0),
	}

}

// AddAttacks can be used for adding more than one attacks
func (bo *Board) AddAttacks(coords []Coord) {

	for _, coord := range coords {
		bo.attacks = append(bo.attacks, coord)
	}
}

func (bo *Board) Score() int {
	score := 0
	for _, ship := range bo.ships {
		for _, attack := range bo.attacks {
			if ship.IsHit(attack.X, attack.Y) {
				score++
			}
		}
	}
	return score
}

func (bo *Board) AddAttack(x, y int) {
	bo.attacks = append(bo.attacks, Coord{x, y})
}

func (bo *Board) Render() [][]string {
	var matrix [][]string

	for i := 0; i < bo.gridSize; i++ {

		matrix = append(matrix, make([]string, bo.gridSize))
		for j := 0; j < bo.gridSize; j++ {
			matrix[i][j] = "_"
		}

	}
	// assume all attacks to miss in the beginning
	for _, attack := range bo.attacks {
		matrix[attack.X][attack.Y] = "O"
	}

	// mark all the hits and boats
	for _, ship := range bo.ships {
		loc := ship.Loc()
		matrix[loc.X][loc.Y] = "B"

		for _, attack := range bo.attacks {
			if ship.IsHit(attack.X, attack.Y) {
				matrix[attack.X][attack.Y] = "X"
			}
		}
	}
	return matrix
}
