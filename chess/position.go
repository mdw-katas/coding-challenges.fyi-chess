package chess

import (
	"strings"

	"github.com/mdwhatcott/must/must"
)

type Position struct {
	White         [6]BitBoard
	Black         [6]BitBoard
	Castling      BitBoard
	EnPassant     BitBoard
	HalfMoveClock uint16
	FullMoveCount uint16
	WhiteToMove   bool
}

func StartingPosition() *Position {
	return must.Value(ParseFEN(startingFEN))
}

func (this *Position) String() string {
	var result strings.Builder
	for s, square := range allSquares {
		length := result.Len()
		for _, pieceType := range allPieceTypes {
			if this.White[pieceType].IsOccupied(square) {
				result.WriteString(pieceType.WhiteFigurine())
			} else if this.Black[pieceType].IsOccupied(square) {
				result.WriteString(pieceType.BlackFigurine())
			}
		}
		if result.Len() == length {
			result.WriteString("-")
		}
		if (s+1)%boardWidth == 0 {
			result.WriteString("\n")
		} else {
			result.WriteString(" ")
		}
	}
	return strings.TrimSpace(result.String())
}
