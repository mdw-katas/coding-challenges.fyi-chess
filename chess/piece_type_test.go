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
	should.So(t, PieceType(42).WhiteFigurine(), should.Equal, "?")

	should.So(t, King.BlackFigurine(), should.Equal, "♚")
	should.So(t, Queen.BlackFigurine(), should.Equal, "♛")
	should.So(t, Rook.BlackFigurine(), should.Equal, "♜")
	should.So(t, Bishop.BlackFigurine(), should.Equal, "♝")
	should.So(t, Knight.BlackFigurine(), should.Equal, "♞")
	should.So(t, Pawn.BlackFigurine(), should.Equal, "♟︎")
	should.So(t, PieceType(42).BlackFigurine(), should.Equal, "¿")
}
