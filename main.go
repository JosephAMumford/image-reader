package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JosephAMumford/image-reader/formats"
)

const PROMPT = "> "

func main() {

	bmpFile := formats.BMP{}

	bmpFile.LoadFile("samples/cat-dragon.bmp")
	bmpFile.Print()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter any single character to display the image. Type 'quit' to quit")

	for {
		fmt.Fprint(os.Stdout, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "quit" {
			break
		}

		if len(line) == 1 {
			bmpFile.Render(line)
		}
	}
}
