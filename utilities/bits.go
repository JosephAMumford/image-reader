package utilities

type nibble struct {
	hi  uint8
	low uint8
}

func NewNibble(i uint8) *nibble {
	nib := nibble{
		hi:  i >> 4 & 0x0f,
		low: i & 0x0f,
	}

	return &nib
}

func GetHiNibble(i uint8) uint8 {
	return i >> 4 & 0x0f
}

func GetLowNibble(i uint8) uint8 {
	return i & 0x0f
}