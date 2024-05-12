package main

import "github.com/JosephAMumford/image-reader/formats"

func main() {

	bmpFile := formats.BMP{}

	bmpFile.LoadFile("samples/simple.bmp")
	bmpFile.Render()
}
