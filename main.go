package main

import (
	"os"

	"github.com/girishramnani/gci17_gojek_challenge/game"
)

func main() {

	game.Play(os.Stdin, os.Stdout)

}
