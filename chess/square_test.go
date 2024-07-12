package chess

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/testing/should"
)

var allSquaresRaw = strings.Join([]string{
	"a8 b8 c8 d8 e8 f8 g8 h8",
	"a7 b7 c7 d7 e7 f7 g7 h7",
	"a6 b6 c6 d6 e6 f6 g6 h6",
	"a5 b5 c5 d5 e5 f5 g5 h5",
	"a4 b4 c4 d4 e4 f4 g4 h4",
	"a3 b3 c3 d3 e3 f3 g3 h3",
	"a2 b2 c2 d2 e2 f2 g2 h2",
	"a1 b1 c1 d1 e1 f1 g1 h1",
}, " ")

func TestSquareRankFile(t *testing.T) {
	should.So(t, a1.File(), should.Equal, 'a')
	should.So(t, b2.File(), should.Equal, 'b')
	should.So(t, c3.File(), should.Equal, 'c')
	should.So(t, d4.File(), should.Equal, 'd')
	should.So(t, e5.File(), should.Equal, 'e')
	should.So(t, f6.File(), should.Equal, 'f')
	should.So(t, g7.File(), should.Equal, 'g')
	should.So(t, h8.File(), should.Equal, 'h')

	should.So(t, a1.Rank(), should.Equal, 1)
	should.So(t, b2.Rank(), should.Equal, 2)
	should.So(t, c3.Rank(), should.Equal, 3)
	should.So(t, d4.Rank(), should.Equal, 4)
	should.So(t, e5.Rank(), should.Equal, 5)
	should.So(t, f6.Rank(), should.Equal, 6)
	should.So(t, g7.Rank(), should.Equal, 7)
	should.So(t, h8.Rank(), should.Equal, 8)
}
func TestParseSquare(t *testing.T) {
	mustParse := func(s string) Square {
		square, err := parseSquare(s)
		should.So(t, err, should.BeNil)
		return square
	}
	should.So(t, funcy.Map(mustParse, strings.Fields(allSquaresRaw)), should.Equal, allSquares)
	should.So(t, strings.Join(funcy.Map(Square.String, allSquares), " "), should.Equal, allSquaresRaw)
}
func TestParseSquareErr(t *testing.T) {
	assertInvalidSquare(t, "")
	assertInvalidSquare(t, "a")
	assertInvalidSquare(t, "aa")
	assertInvalidSquare(t, "1")
	assertInvalidSquare(t, "11")
	assertInvalidSquare(t, "a9")
	assertInvalidSquare(t, "a0")
	assertInvalidSquare(t, "i1")
	assertInvalidSquare(t, "a12")
}
func assertInvalidSquare(t *testing.T, input string) {
	t.Run(input, func(t *testing.T) {
		parsed, err := parseSquare(input)
		should.So(t, parsed, should.Equal, 0)
		should.So(t, err, should.WrapError, errInvalidSquare)
	})
}
