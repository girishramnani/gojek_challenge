package game

type Coord struct {
	X int
	Y int
}

type Ship struct {
	location Coord
}

func Ships(coords []Coord) []Ship {
	var ships []Ship

	for _, coord := range coords {
		ships = append(ships, Ship{coord})
	}

	return ships
}

func (sh Ship) Loc() Coord {
	return sh.location
}
func (sh Ship) IsHit(x, y int) bool {
	if sh.location.X == x && sh.location.Y == y {
		return true
	}
	return false
}
