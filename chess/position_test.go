package chess

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestPositionSuite(t *testing.T) {
	should.Run(&PositionSuite{T: should.New(t)}, should.Options.UnitTests())
}

type PositionSuite struct {
	*should.T
}

func (this *PositionSuite) Setup() {
}

func (this *PositionSuite) TestString() {
	this.So(new(Position).String(), should.Equal, Blank)
	this.So(StartingPosition().String(), should.Equal, Starting)
}
func (this *PositionSuite) TestParseFEN_InvalidFieldCounts() {
	this.assertInvalidFEN("1 2 3 4 5")     // too few fields  (5)
	this.assertInvalidFEN("1 2 3 4 5 6 7") // too many fields (7)
}
func (this *PositionSuite) TestParseFEN_InvalidRankCounts() {
	this.assertInvalidFEN("8/8/8/8/8/8/8 b KQkq - 0 1")     // too few ranks  (7)
	this.assertInvalidFEN("8/8/8/8/8/8/8/8/8 b KQkq - 0 1") // too many ranks (9)
}
func (this *PositionSuite) TestParseFEN_InvalidRank() {
	this.assertInvalidFEN("8/9/8/8/8/8/8/8 b KQkq - 0 1") // invalid digit '9'
	this.assertInvalidFEN("8/a/8/8/8/8/8/8 b KQkq - 0 1") // invalid character 'a'
}
func (this *PositionSuite) assertInvalidFEN(rawFEN string) {
	position, err := ParseFEN(rawFEN)
	this.So(position, should.BeNil)
	this.So(err, should.WrapError, errInvalidFEN)
	this.Println(err)
}

var (
	Blank = lines(
		"- - - - - - - -",
		"- - - - - - - -",
		"- - - - - - - -",
		"- - - - - - - -",
		"- - - - - - - -",
		"- - - - - - - -",
		"- - - - - - - -",
		"- - - - - - - -",
	)
	Starting = lines(
		"♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜",
		"♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟",
		"- - - - - - - -",
		"- - - - - - - -",
		"- - - - - - - -",
		"- - - - - - - -",
		"♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙",
		"♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖",
	)
	RuyLopezFEN = "r1bqkbnr/pppp1ppp/2n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R b KQkq - 3 3"
	RuyLopez    = lines(
		"♜ - ♝ ♛ ♚ ♝ ♞ ♜",
		"♟ ♟ ♟ ♟ - ♟ ♟ ♟",
		"- - ♞ - - - - -",
		"- ♗ - - ♟ - - -",
		"- - - - ♙ - - -",
		"- - - - - ♘ - -",
		"♙ ♙ ♙ ♙ - ♙ ♙ ♙",
		"♖ ♘ ♗ ♕ ♔ - - ♖",
	)
	MiddleGameFEN = "r1k4r/p2nb1p1/2b4p/1p1n1p2/2PP4/3Q1NB1/1P3PPP/R5K1 w - - 10 20"
	MiddleGame    = lines(
		"♜ - ♚ - - - - ♜",
		"♟ - - ♞ ♝ - ♟ -",
		"- - ♝ - - - - ♟",
		"- ♟ - ♞ - ♟ - -",
		"- - ♙ ♙ - - - -",
		"- - - ♕ - ♘ ♗ -",
		"- ♙ - - - ♙ ♙ ♙",
		"♖ - - - - - ♔ -",
	)
)

func lines(values ...string) string {
	return strings.Join(values, "\n")
}
