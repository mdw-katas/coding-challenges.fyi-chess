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

func (this *Position) Do(move Move) (that *Position) {
	that = &Position{
		Pieces:        this.Pieces,
		Castling:      this.Castling,
		EnPassant:     this.EnPassant,
		HalfMoveClock: this.HalfMoveClock + 1,
		FullMoveCount: this.FullMoveCount,
		ToMove:        !this.ToMove,
	}
	if move.Color == Black {
		that.FullMoveCount++
	}
	for color, pieces := range that.Pieces {
		for piece := range pieces {
			if color == move.Color.Int() && move.Piece == Piece(piece) {
				that.Pieces[color][piece].Vacate(move.Source)
				that.Pieces[color][piece].Occupy(move.Target)
			}
		}
	}
	return that
}
