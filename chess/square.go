package chess

import (
	"errors"
	"fmt"
)

var allSquares = []Square{
	a8, b8, c8, d8, e8, f8, g8, h8,
	a7, b7, c7, d7, e7, f7, g7, h7,
	a6, b6, c6, d6, e6, f6, g6, h6,
	a5, b5, c5, d5, e5, f5, g5, h5,
	a4, b4, c4, d4, e4, f4, g4, h4,
	a3, b3, c3, d3, e3, f3, g3, h3,
	a2, b2, c2, d2, e2, f2, g2, h2,
	a1, b1, c1, d1, e1, f1, g1, h1,
}

const (
	whiteKingsideCastleTarget  = g1
	whiteQueensideCastleTarget = c1

	blackKingsideCastleTarget  = g8
	blackQueensideCastleTarget = c8
)

const (
	minFile = 'a'
	minRank = '1'
)

const boardWidth = 8

func parseSquare(square string) (Square, error) {
	if len(square) != 2 {
		return 0, fmt.Errorf("%w: %s", errInvalidSquare, square)
	}
	switch square[0] {
	case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h':
	default:
		return 0, fmt.Errorf("%w: %s", errInvalidSquare, square)
	}
	switch square[1] {
	case '1', '2', '3', '4', '5', '6', '7', '8':
	default:
		return 0, fmt.Errorf("%w: %s", errInvalidSquare, square)
	}
	file := square[0] - minFile
	rank := square[1] - minRank
	return Square(file + (boardWidth * rank)), nil
}

func (this Square) String() string {
	return fmt.Sprintf("%c%d", this.File(), this.Rank())
}
func (this Square) File() rune {
	return minFile + rune(this%boardWidth)
}
func (this Square) Rank() int {
	return int(this/boardWidth) + 1
}

var errInvalidSquare = errors.New("invalid square")

type Square uint8

const (
	a1 Square = iota
	b1
	c1
	d1
	e1
	f1
	g1
	h1
	a2
	b2
	c2
	d2
	e2
	f2
	g2
	h2
	a3
	b3
	c3
	d3
	e3
	f3
	g3
	h3
	a4
	b4
	c4
	d4
	e4
	f4
	g4
	h4
	a5
	b5
	c5
	d5
	e5
	f5
	g5
	h5
	a6
	b6
	c6
	d6
	e6
	f6
	g6
	h6
	a7
	b7
	c7
	d7
	e7
	f7
	g7
	h7
	a8
	b8
	c8
	d8
	e8
	f8
	g8
	h8
)
