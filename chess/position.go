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
	for _, color := range []Color{White, Black} {
		for _, piece := range allPieceTypes {
			if move.Color == color && move.Piece == piece {
				nColor := color.Int()
				that.Pieces[nColor][piece].Vacate(move.Source)
				that.Pieces[nColor][piece].Occupy(move.Target)
			}
		}
	}
	return that
}

func StartingPosition() *Position {
	return must.Value(ParseFEN(startingFEN))
}
