package main

import (
	"strings"
)

// Commonly used bit array data structure for board games.
// Each bit in the 64-bit integer corresponds with a chess 
// board square. Conventionally, bits are associated with
// squares from left to right, bottom to top. bit 0 (LSB)
// is A1, bit 7 is H1, bit 56 is A8 and bit 63 is H8. 
type Bitboard uint64

// Create a new bitboard from a map where m[Square] = true
// if the square is occupied.
func NewBitboard(m map[Square]bool) Bitboard {
	var bb uint64
	for sq := 63; sq >= 0; sq-- {
		bb <<= 1
		if m[Square(sq)] {
			bb |= 1
		}
	}
	return Bitboard(bb)
}

// Return a string representation of the bitboard.
func (bb Bitboard) Draw() string {
	var sb strings.Builder
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			if uint64(bb) & (1 << uint8(rank*8+file)) != 0 {
				// occupied
				sb.WriteString("X ")
			} else {
				// unoccupied
				sb.WriteString(". ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
