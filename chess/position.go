package chess

import "github.com/mdwhatcott/must/must"

type Position struct {
	Pieces        [2][6]BitBoard
	Castling      BitBoard
	EnPassant     BitBoard
	HalfMoveClock uint16
	FullMoveCount uint16
	ToMove        Color
}

func StartingPosition() *Position {
	return must.Value(ParseFEN(startingFEN))
}
