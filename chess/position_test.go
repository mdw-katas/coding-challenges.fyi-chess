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
func (this *PositionSuite) TestParseFEN_InvalidPlayerToMove() {
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 ? KQkq - 0 1") // invalid player '?'
}
func (this *PositionSuite) TestParseFEN_ValidPlayerToMove() {
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - - 0 1").WhiteToMove, should.BeTrue)
	this.So(this.mustParse("8/8/8/8/8/8/8/8 b - - 0 1").WhiteToMove, should.BeFalse)
}
func (this *PositionSuite) TestParseFEN_InvalidCastlingRights() {
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 w asdf - 0 1") // invalid castling 'asdf'
}
func (this *PositionSuite) TestParseFEN_Castling_WhiteKingside() {
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w K - 0 1").Castling.IsOccupied(whiteKingsideCastleTarget), should.BeTrue)
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w Q - 0 1").Castling.IsOccupied(whiteQueensideCastleTarget), should.BeTrue)
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w k - 0 1").Castling.IsOccupied(blackKingsideCastleTarget), should.BeTrue)
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w q - 0 1").Castling.IsOccupied(blackQueensideCastleTarget), should.BeTrue)
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w KQkq - 0 1").Castling, should.Equal, *NewBitBoard(
		whiteKingsideCastleTarget,
		whiteQueensideCastleTarget,
		blackKingsideCastleTarget,
		blackQueensideCastleTarget,
	))
}
func (this *PositionSuite) mustParse(rawFEN string) *Position {
	position, err := ParseFEN(rawFEN)
	this.So(err, should.BeNil)
	return position
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
