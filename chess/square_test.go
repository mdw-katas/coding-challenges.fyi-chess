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
