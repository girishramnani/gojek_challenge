package game

import "testing"

func TestShipIsHitTrue(t *testing.T) {

	ship := Ship{Coord{4, 5}}

	if !ship.IsHit(4, 5) {
		t.Error("IsHit expected to return true but returned false")
	}

}
func TestShipIsHitFalse(t *testing.T) {

	ship := Ship{Coord{5, 7}}

	if ship.IsHit(1, 2) {
		t.Error("IsHit expected to return false but returned true")
	}

}
