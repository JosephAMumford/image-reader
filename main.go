package main

import (
	"github.com/JosephAMumford/image-reader/formats"
)

func main() {

	bmpFile := formats.BMP{}

	bmpFile.LoadFile("samples/cat-dragon.bmp")
	bmpFile.Render("\u2588")
}
