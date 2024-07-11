package chess

import (
	"testing"

	"github.com/mdwhatcott/must/must"
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
	this.So(must.Value(ParseFEN(RuyLopezFEN)).String(), should.Equal, RuyLopez)
	this.So(must.Value(ParseFEN(MiddleGameFEN)).String(), should.Equal, MiddleGame)
}
