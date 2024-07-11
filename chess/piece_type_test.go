package chess

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestPieceTypeFigurines(t *testing.T) {
	should.So(t, King.Figurine(White), should.Equal, "♔")
	should.So(t, Queen.Figurine(White), should.Equal, "♕")
	should.So(t, Rook.Figurine(White), should.Equal, "♖")
	should.So(t, Bishop.Figurine(White), should.Equal, "♗")
	should.So(t, Knight.Figurine(White), should.Equal, "♘")
	should.So(t, Pawn.Figurine(White), should.Equal, "♙")
	should.So(t, Piece(42).Figurine(White), should.Equal, "?")

	should.So(t, King.Figurine(Black), should.Equal, "♚")
	should.So(t, Queen.Figurine(Black), should.Equal, "♛")
	should.So(t, Rook.Figurine(Black), should.Equal, "♜")
	should.So(t, Bishop.Figurine(Black), should.Equal, "♝")
	should.So(t, Knight.Figurine(Black), should.Equal, "♞")
	should.So(t, Pawn.Figurine(Black), should.Equal, "♟")
	should.So(t, Piece(42).Figurine(Black), should.Equal, "¿")
}
func TestPieceTypeInitials(t *testing.T) {
	should.So(t, King.Initial(White), should.Equal, "K")
	should.So(t, Queen.Initial(White), should.Equal, "Q")
	should.So(t, Rook.Initial(White), should.Equal, "R")
	should.So(t, Bishop.Initial(White), should.Equal, "B")
	should.So(t, Knight.Initial(White), should.Equal, "N")
	should.So(t, Pawn.Initial(White), should.Equal, "P")
	should.So(t, Piece(42).Initial(White), should.Equal, "?")

	should.So(t, King.Initial(Black), should.Equal, "k")
	should.So(t, Queen.Initial(Black), should.Equal, "q")
	should.So(t, Rook.Initial(Black), should.Equal, "r")
	should.So(t, Bishop.Initial(Black), should.Equal, "b")
	should.So(t, Knight.Initial(Black), should.Equal, "n")
	should.So(t, Pawn.Initial(Black), should.Equal, "p")
	should.So(t, Piece(42).Initial(Black), should.Equal, "¿")
}
