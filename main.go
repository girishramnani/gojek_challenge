package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/girishramnani/gci17_gojek_challenge/game"
)

func main() {

	input := flag.String("input", "", "The input file")
	output := flag.String("output", "", "The output file")

	flag.Parse()

	inputStream := os.Stdin
	outputStream := os.Stdout

	if *input != "" {
		f, err := os.Open(*input)
		if err != nil {
			fmt.Println("Cannot open the file ", *input)
			os.Exit(-1)
		}
		inputStream = f
	}
	if *output != "" {
		f, err := os.Create(*output)
		if err != nil {
			fmt.Println("Cannot create the file ", *output)
			os.Exit(-1)
		}
		outputStream = f
	}

	game.Play(inputStream, outputStream)

}
