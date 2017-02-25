package game

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Play(reader io.Reader, writer io.Writer) {

	var gridSize, numShips int
	var numAttacks int

	// Input
	fmt.Fscanf(reader, "%d", &gridSize)
	fmt.Fscanf(reader, "%d", &numShips)
	pOneCoords := CoordInput(reader, ",", ":")
	pTwoCoords := CoordInput(reader, ",", ":")
	fmt.Fscanf(reader, "%d", &numAttacks)
	attackOne := CoordInput(reader, ":", ",")
	attackTwo := CoordInput(reader, ":", ",")

	boardOne := NewBoard(gridSize, Ships(pOneCoords)...)
	boardTwo := NewBoard(gridSize, Ships(pTwoCoords)...)
	boardOne.AddAttacks(attackTwo)
	boardTwo.AddAttacks(attackOne)
	scoreOne := boardOne.Score()
	scoreTwo := boardTwo.Score()

	// Output
	fmt.Fprintln(writer, "Player1")
	for _, line := range boardOne.Render() {
		fmt.Fprintln(writer, strings.Join(line, " "))
	}
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Player2")
	for _, line := range boardTwo.Render() {
		fmt.Fprintln(writer, strings.Join(line, " "))
	}
	fmt.Fprintln(writer)

	fmt.Fprintf(writer, "P1:%d\n", scoreOne)
	fmt.Fprintf(writer, "P2:%d\n", scoreTwo)

	switch {
	case scoreOne < scoreTwo:
		fmt.Fprintln(writer, "Player 2 wins")
	case scoreOne > scoreTwo:
		fmt.Fprintln(writer, "Player 1 wins")
	default:
		fmt.Fprintln(writer, "It is a draw")
	}

}

func CoordInput(reader io.Reader, linesep string, coordsep string) []Coord {

	var coords []Coord

	var line string
	fmt.Fscan(reader, &line)

	strcoords := strings.Split(line, linesep)

	var temp []string
	var x, y int
	for _, strc := range strcoords {
		temp = strings.Split(strc, coordsep)
		x, _ = strconv.Atoi(temp[0])
		y, _ = strconv.Atoi(temp[1])
		coords = append(coords, Coord{X: x, Y: y})
	}
	return coords

}
