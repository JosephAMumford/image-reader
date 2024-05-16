package formats

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/gookit/color"
)

type BMP struct {
	Id                      string
	Filesize                uint32
	Reserved                uint32
	Offset                  uint32
	BitMapInfoHeader        uint32
	HorizontalWidth         uint32
	VerticalWidth           uint32
	NumberOfPlanes          uint16
	BitsPerPixel            uint16
	CompressionType         uint32
	SizeInBytes             uint32
	HorizontalResolution    uint32
	VerticalResolution      uint32
	NumberOfUsedColors      uint32
	NumberOfImportantColors uint32
	PixelData               []byte
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (b *BMP) LoadFile(filename string) {
	f, err := os.Open(filename)
	check(err)

	defer f.Close()

	currentOffset := 0

	//Id
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	id := make([]byte, 2)
	bytesRead, err := f.Read(id)
	check(err)
	b.Id = string(id)

	if b.Id != "BM" && b.Id != "BA" && b.Id != "CI" && b.Id != "CP" && b.Id != "IC" && b.Id != "PT" {
		fmt.Println("Not a BMP file")
		return
	}

	currentOffset += bytesRead

	//Filesize
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	fileSize := make([]byte, 4)
	bytesRead, err = f.Read(fileSize)
	check(err)
	b.Filesize = binary.LittleEndian.Uint32(fileSize)
	currentOffset += bytesRead

	//Reserved
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	reserved := make([]byte, 4)
	bytesRead, err = f.Read(reserved)
	check(err)
	b.Reserved = binary.LittleEndian.Uint32(reserved)
	currentOffset += bytesRead

	//Offset
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	offset := make([]byte, 4)
	bytesRead, err = f.Read(offset)
	check(err)
	b.Offset = binary.LittleEndian.Uint32(offset)
	currentOffset += bytesRead

	//BitMapInfoHeader
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	bitMapInfoHeader := make([]byte, 4)
	bytesRead, err = f.Read(bitMapInfoHeader)
	check(err)
	b.BitMapInfoHeader = binary.LittleEndian.Uint32(bitMapInfoHeader)
	currentOffset += bytesRead

	//Horizontal Width
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	horizontalWidth := make([]byte, 4)
	bytesRead, err = f.Read(horizontalWidth)
	check(err)
	b.HorizontalWidth = binary.LittleEndian.Uint32(horizontalWidth)
	currentOffset += bytesRead

	//Vertical Width
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	verticalWidth := make([]byte, 4)
	bytesRead, err = f.Read(verticalWidth)
	check(err)
	b.VerticalWidth = binary.LittleEndian.Uint32(verticalWidth)
	currentOffset += bytesRead

	//Number of Planes
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	numberOfPlanes := make([]byte, 2)
	bytesRead, err = f.Read(numberOfPlanes)
	check(err)
	b.NumberOfPlanes = binary.LittleEndian.Uint16(numberOfPlanes)
	currentOffset += bytesRead

	//Bits per pixel
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	bitsPerPixel := make([]byte, 2)
	bytesRead, err = f.Read(bitsPerPixel)
	check(err)
	b.BitsPerPixel = binary.LittleEndian.Uint16(bitsPerPixel)
	currentOffset += bytesRead

	//Compression Type
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	compressionType := make([]byte, 4)
	bytesRead, err = f.Read(compressionType)
	check(err)
	b.CompressionType = binary.LittleEndian.Uint32(compressionType)
	currentOffset += bytesRead

	//Size in Bytes
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	sizeInBytes := make([]byte, 4)
	bytesRead, err = f.Read(sizeInBytes)
	check(err)
	b.SizeInBytes = binary.LittleEndian.Uint32(sizeInBytes)
	currentOffset += bytesRead

	//Horizontal Resolution
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	horizontalResolution := make([]byte, 4)
	bytesRead, err = f.Read(horizontalResolution)
	check(err)
	b.HorizontalResolution = binary.LittleEndian.Uint32(horizontalResolution)
	currentOffset += bytesRead

	//Vertical Resolution
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	verticalResolution := make([]byte, 4)
	bytesRead, err = f.Read(verticalResolution)
	check(err)
	b.VerticalResolution = binary.LittleEndian.Uint32(verticalResolution)
	currentOffset += bytesRead

	//Number of Used Colors
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	numberOfUsedColors := make([]byte, 4)
	bytesRead, err = f.Read(numberOfUsedColors)
	check(err)
	b.NumberOfUsedColors = binary.LittleEndian.Uint32(numberOfUsedColors)
	currentOffset += bytesRead

	//Number of Important Colors
	_, err = f.Seek(int64(currentOffset), io.SeekStart)
	check(err)
	numberOfImportantColors := make([]byte, 4)
	bytesRead, err = f.Read(numberOfImportantColors)
	check(err)
	b.NumberOfImportantColors = binary.LittleEndian.Uint32(numberOfImportantColors)
	currentOffset += bytesRead

	//Pixel Data
	b.PixelData = make([]byte, 3 * b.HorizontalWidth * b.VerticalWidth)
	_, err = f.Seek(int64(b.Offset), io.SeekStart)
	check(err)
	_, err = f.Read(b.PixelData)
	check(err)
}

// Use single char as string
func (b *BMP) Render(char string) {
	//Y starts at the end because BMP data is load bottom to top, left to right
	for y := int(b.VerticalWidth) - 1; y > -1 ; y-- {
		for x := 0; x < int(b.HorizontalWidth); x++ {
			index := y * (int(b.HorizontalWidth) * 3) + (x * 3)
			color.RGB(b.PixelData[index + 2],b.PixelData[index + 1],b.PixelData[index]).Print(char)
		}
		fmt.Print("\n")
	}
}

func (b * BMP) Print() {
	fmt.Printf("Id: %s \n", b.Id)
	fmt.Printf("Filesize: %d \n", b.Filesize)
	fmt.Printf("Offset: %d \n", b.Offset)
	fmt.Printf("BitMapInfoHeader: %d \n", b.BitMapInfoHeader)
	fmt.Printf("Horizontal Width: %d \n", b.HorizontalWidth)
	fmt.Printf("Vertical Width: %d \n", b.VerticalWidth)
	fmt.Printf("Number of Planes: %d \n", b.NumberOfPlanes)
	fmt.Printf("Bits Per Pixel: %d \n", b.BitsPerPixel)
	fmt.Printf("Compression Type: %d \n", b.CompressionType)
	fmt.Printf("Size In bytes: %d \n", b.SizeInBytes)
	fmt.Printf("Horizontal Resolution: %d \n", b.HorizontalResolution)
	fmt.Printf("Vertical Resolution: %d \n", b.VerticalResolution)
	fmt.Printf("Number of Used Colors: %d \n", b.NumberOfUsedColors)
	fmt.Printf("Number Of Important Colors: %d \n", b.NumberOfImportantColors)
}