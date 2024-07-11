package chess

import "github.com/mdwhatcott/must/must"

type Position struct {
	WhitePieces   [6]BitBoard
	BlackPieces   [6]BitBoard
	Castling      BitBoard
	EnPassant     BitBoard
	HalfMoveClock uint16
	FullMoveCount uint16
	ToMove        Color
}

func StartingPosition() *Position {
	return must.Value(ParseFEN(startingFEN))
}
