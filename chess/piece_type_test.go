package chess

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestPieceTypeFigurines(t *testing.T) {
	should.So(t, King.WhiteFigurine(), should.Equal, "♔")
	should.So(t, Queen.WhiteFigurine(), should.Equal, "♕")
	should.So(t, Rook.WhiteFigurine(), should.Equal, "♖")
	should.So(t, Bishop.WhiteFigurine(), should.Equal, "♗")
	should.So(t, Knight.WhiteFigurine(), should.Equal, "♘")
	should.So(t, Pawn.WhiteFigurine(), should.Equal, "♙")
	should.So(t, Piece(42).WhiteFigurine(), should.Equal, "?")

	should.So(t, King.BlackFigurine(), should.Equal, "♚")
	should.So(t, Queen.BlackFigurine(), should.Equal, "♛")
	should.So(t, Rook.BlackFigurine(), should.Equal, "♜")
	should.So(t, Bishop.BlackFigurine(), should.Equal, "♝")
	should.So(t, Knight.BlackFigurine(), should.Equal, "♞")
	should.So(t, Pawn.BlackFigurine(), should.Equal, "♟")
	should.So(t, Piece(42).BlackFigurine(), should.Equal, "¿")
}
func TestPieceTypeInitials(t *testing.T) {
	should.So(t, King.WhiteInitial(), should.Equal, "K")
	should.So(t, Queen.WhiteInitial(), should.Equal, "Q")
	should.So(t, Rook.WhiteInitial(), should.Equal, "R")
	should.So(t, Bishop.WhiteInitial(), should.Equal, "B")
	should.So(t, Knight.WhiteInitial(), should.Equal, "N")
	should.So(t, Pawn.WhiteInitial(), should.Equal, "P")
	should.So(t, Piece(42).WhiteInitial(), should.Equal, "?")

	should.So(t, King.BlackInitial(), should.Equal, "k")
	should.So(t, Queen.BlackInitial(), should.Equal, "q")
	should.So(t, Rook.BlackInitial(), should.Equal, "r")
	should.So(t, Bishop.BlackInitial(), should.Equal, "b")
	should.So(t, Knight.BlackInitial(), should.Equal, "n")
	should.So(t, Pawn.BlackInitial(), should.Equal, "p")
	should.So(t, Piece(42).BlackInitial(), should.Equal, "¿")
}
