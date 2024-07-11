package chess

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestFENSuite(t *testing.T) {
	should.Run(&FENSuite{T: should.New(t)}, should.Options.UnitTests())
}

type FENSuite struct {
	*should.T
}

func (this *FENSuite) TestParseFEN_InvalidFieldCounts() {
	this.assertInvalidFEN("1 2 3 4 5")     // too few fields  (5)
	this.assertInvalidFEN("1 2 3 4 5 6 7") // too many fields (7)
}
func (this *FENSuite) TestParseFEN_InvalidRankCounts() {
	this.assertInvalidFEN("8/8/8/8/8/8/8 b KQkq - 0 1")     // too few ranks  (7)
	this.assertInvalidFEN("8/8/8/8/8/8/8/8/8 b KQkq - 0 1") // too many ranks (9)
}
func (this *FENSuite) TestParseFEN_InvalidRank() {
	this.assertInvalidFEN("8/9/8/8/8/8/8/8 b KQkq - 0 1") // invalid digit '9'
	this.assertInvalidFEN("8/a/8/8/8/8/8/8 b KQkq - 0 1") // invalid character 'a'
}
func (this *FENSuite) TestParseFEN_InvalidPlayerToMove() {
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 ? KQkq - 0 1") // invalid player '?'
}
func (this *FENSuite) TestParseFEN_ValidPlayerToMove() {
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - - 0 1").WhiteToMove, should.BeTrue)
	this.So(this.mustParse("8/8/8/8/8/8/8/8 b - - 0 1").WhiteToMove, should.BeFalse)
}
func (this *FENSuite) TestParseFEN_InvalidCastlingRights() {
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 w asdf - 0 1") // invalid castling 'asdf'
}
func (this *FENSuite) TestParseFEN_Castling_WhiteKingside() {
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
func (this *FENSuite) TestParseFEN_InvalidEnPassantTarget() {
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 w - a 0 1")  // 'a' isn't complete
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 w - 1a 0 1") // '1a' is backwards
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 w - a0 0 1") // 'a9' isn't possible
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 w - a1 0 1") // 'a1' isn't possible
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 w - a2 0 1") // 'a2' isn't possible
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 w - a4 0 1") // 'a4' isn't possible
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 w - a5 0 1") // 'a5' isn't possible
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 w - a7 0 1") // 'a7' isn't possible
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 w - a8 0 1") // 'a8' isn't possible
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 w - a9 0 1") // 'a9' isn't possible
}
func (this *FENSuite) TestParseFEN_ValidEnPassantTarget() {
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - a3 0 1").EnPassant, should.Equal, *NewBitBoard(a3))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - a6 0 1").EnPassant, should.Equal, *NewBitBoard(a6))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - b3 0 1").EnPassant, should.Equal, *NewBitBoard(b3))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - b6 0 1").EnPassant, should.Equal, *NewBitBoard(b6))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - c3 0 1").EnPassant, should.Equal, *NewBitBoard(c3))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - c6 0 1").EnPassant, should.Equal, *NewBitBoard(c6))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - d3 0 1").EnPassant, should.Equal, *NewBitBoard(d3))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - d6 0 1").EnPassant, should.Equal, *NewBitBoard(d6))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - e3 0 1").EnPassant, should.Equal, *NewBitBoard(e3))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - e6 0 1").EnPassant, should.Equal, *NewBitBoard(e6))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - f3 0 1").EnPassant, should.Equal, *NewBitBoard(f3))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - f6 0 1").EnPassant, should.Equal, *NewBitBoard(f6))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - g3 0 1").EnPassant, should.Equal, *NewBitBoard(g3))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - g6 0 1").EnPassant, should.Equal, *NewBitBoard(g6))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - h3 0 1").EnPassant, should.Equal, *NewBitBoard(h3))
	this.So(this.mustParse("8/8/8/8/8/8/8/8 w - h6 0 1").EnPassant, should.Equal, *NewBitBoard(h6))
}
func (this *FENSuite) TestParseFEN_InvalidHalfMoveClock() {
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 b - - -1 1") // can't be negative
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 b - - 01 1") // can't be zero-prefixed
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 b - - a 1")  // must be a number
}
func (this *FENSuite) TestParseFEN_ValidHalfMoveClock() {
	this.So(this.mustParse("8/8/8/8/8/8/8/8 b - - 0 1").HalfMoveClock, should.Equal, 0)
	this.So(this.mustParse("8/8/8/8/8/8/8/8 b - - 1 1").HalfMoveClock, should.Equal, 1)
	this.So(this.mustParse("8/8/8/8/8/8/8/8 b - - 50 1").HalfMoveClock, should.Equal, 50)
}
func (this *FENSuite) TestParseFEN_InvalidFullMoveCount() {
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 b - - 0 0")  // can't be zero
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 b - - 0 -1") // can't be negative
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 b - - 0 01") // can't be zero-prefixed
	this.assertInvalidFEN("8/8/8/8/8/8/8/8 b - - 0 a")  // must be a number
}
func (this *FENSuite) TestParseFEN_ValidFullMoveCount() {
	this.So(this.mustParse("8/8/8/8/8/8/8/8 b - - 0 1").FullMoveCount, should.Equal, 1)
	this.So(this.mustParse("8/8/8/8/8/8/8/8 b - - 0 50").FullMoveCount, should.Equal, 50)
}
func (this *FENSuite) mustParse(rawFEN string) *Position {
	position, err := ParseFEN(rawFEN)
	this.So(err, should.BeNil)
	return position
}
func (this *FENSuite) assertInvalidFEN(rawFEN string) {
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
