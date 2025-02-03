package main

import "strings"

// A chess board position describes where all of the pieces
// are placed as well as game state such as whose turn it is,
// castling rights, and the number of moves played. In short,
// a position should include all of the game state.
type Position struct {
	WhitePawnBB   Bitboard
	WhiteBishopBB Bitboard
	WhiteKnightBB Bitboard
	WhiteRookBB   Bitboard
	WhiteQueenBB  Bitboard
	WhiteKingBB   Bitboard
	BlackPawnBB   Bitboard
	BlackBishopBB Bitboard
	BlackKnightBB Bitboard
	BlackRookBB   Bitboard
	BlackQueenBB  Bitboard
	BlackKingBB   Bitboard
}

// Create starting position
func NewStartPosition() Position {
	var p Position
	p.WhitePawnBB = NewBitboard(map[Square]bool{A2: true, B2: true, C2: true, D2: true, E2: true, F2: true, G2: true, H2: true})
	p.WhiteBishopBB = NewBitboard(map[Square]bool{C1: true, F1: true})
	p.WhiteKnightBB = NewBitboard(map[Square]bool{B1: true, G1: true})
	p.WhiteRookBB = NewBitboard(map[Square]bool{A1: true, H1: true})
	p.WhiteQueenBB = NewBitboard(map[Square]bool{D1: true})
	p.WhiteKingBB = NewBitboard(map[Square]bool{E1: true})

	p.BlackPawnBB = NewBitboard(map[Square]bool{A7: true, B7: true, C7: true, D7: true, E7: true, F7: true, G7: true, H7: true})
	p.BlackBishopBB = NewBitboard(map[Square]bool{C8: true, F8: true})
	p.BlackKnightBB = NewBitboard(map[Square]bool{B8: true, G8: true})
	p.BlackRookBB = NewBitboard(map[Square]bool{A8: true, H8: true})
	p.BlackQueenBB = NewBitboard(map[Square]bool{D8: true})
	p.BlackKingBB = NewBitboard(map[Square]bool{E8: true})

	return p
}

// Create position from FEN
func NewPositionFromFEN(FEN string) {}

func (p Position) String() string {
	var sb strings.Builder
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			sq := uint8(rank*8 + file)
			mask := uint64(1 << sq)
			switch {
			case uint64(p.WhitePawnBB)&mask != 0:
				sb.WriteString("P ")
			case uint64(p.WhiteBishopBB)&mask != 0:
				sb.WriteString("B ")
			case uint64(p.WhiteKnightBB)&mask != 0:
				sb.WriteString("N ")
			case uint64(p.WhiteRookBB)&mask != 0:
				sb.WriteString("R ")
			case uint64(p.WhiteQueenBB)&mask != 0:
				sb.WriteString("Q ")
			case uint64(p.WhiteKingBB)&mask != 0:
				sb.WriteString("K ")
			case uint64(p.BlackPawnBB)&mask != 0:
				sb.WriteString("p ")
			case uint64(p.BlackBishopBB)&mask != 0:
				sb.WriteString("b ")
			case uint64(p.BlackKnightBB)&mask != 0:
				sb.WriteString("n ")
			case uint64(p.BlackRookBB)&mask != 0:
				sb.WriteString("r ")
			case uint64(p.BlackQueenBB)&mask != 0:
				sb.WriteString("q ")
			case uint64(p.BlackKingBB)&mask != 0:
				sb.WriteString("k ")
			default:
				sb.WriteString(". ") // Empty square
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

// TODO: MOVE LATER

type Piece int8

const (
	WhitePawn   = 1
	WhiteBishop = 2
	WhiteKnight = 3
	WhiteRook   = 4
	WhiteQueen  = 5
	WhiteKing   = 6
	BlackPawn   = -1
	BlackBishop = -2
	BlackKnight = -3
	BlackRook   = -4
	BlackQueen  = -5
	BlackKing   = -6
)
