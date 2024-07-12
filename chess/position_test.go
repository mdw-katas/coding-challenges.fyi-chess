package chess

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestPositionSuite(t *testing.T) {
	should.Run(&PositionSuite{T: should.New(t)}, should.Options.UnitTests())
}

type PositionSuite struct {
	*should.T
	position *Position
}

func (this *PositionSuite) Setup() {
	this.position = StartingPosition()
}

func (this *PositionSuite) TestDoMoves() {
	moves := []Move{
		{Color: White, Piece: Pawn, Source: f2, Target: f3},
		{Color: Black, Piece: Pawn, Source: e7, Target: e6},
		{Color: White, Piece: Pawn, Source: g2, Target: g4},
		{Color: Black, Piece: Queen, Source: d8, Target: h4},
	}
	for _, move := range moves {
		this.position = this.position.Do(move)
	}
	this.So(PlainText(this.position), should.Equal, FoolsMate)
}

var FoolsMate = lines(
	"♜ ♞ ♝ - ♚ ♝ ♞ ♜",
	"♟ ♟ ♟ ♟ - ♟ ♟ ♟",
	"- - - - ♟ - - -",
	"- - - - - - - -",
	"- - - - - - ♙ ♛",
	"- - - - - ♙ - -",
	"♙ ♙ ♙ ♙ ♙ - - ♙",
	"♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖",
)
