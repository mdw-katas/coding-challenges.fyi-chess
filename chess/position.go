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

func (this *Position) Place(color Color, piece Piece, square Square) {
	this.Pieces[color.Int()][piece].Occupy(square)
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

func GenerateMoves(position *Position) func(func(Move) bool) {
	return func(yield func(Move) bool) {
		color := position.ToMove
		for pieceN, board := range position.Pieces[color.Int()] {
			for square := range board.OccupiedSquares() {
				switch Piece(pieceN) {
				case Pawn:
					if color == White {
						if square.Rank() == 7 {
							yield(newPromotion(newMove(color, Pawn, square, square+8), Queen))
							yield(newPromotion(newMove(color, Pawn, square, square+8), Rook))
							yield(newPromotion(newMove(color, Pawn, square, square+8), Bishop))
							yield(newPromotion(newMove(color, Pawn, square, square+8), Knight))
						} else {
							yield(newMove(color, Pawn, square, square+8))
						}
						if square.Rank() == 2 {
							yield(newMove(color, Pawn, square, square+16))
						}
					}
				}
			}
		}
	}
}

func newMove(color Color, piece Piece, source, target Square) Move {
	return Move{Color: color, Piece: piece, Source: source, Target: target}
}
func newPromotion(move Move, promotion Piece) Move {
	move.Promotion = promotion
	return move
}
