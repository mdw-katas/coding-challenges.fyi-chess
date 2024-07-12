package chess

import (
	"testing"

	"github.com/mdwhatcott/must/must"
	"github.com/mdwhatcott/testing/should"
)

func TestTUISuite(t *testing.T) {
	should.Run(&TUISuite{T: should.New(t)}, should.Options.UnitTests())
}

type TUISuite struct {
	*should.T
}

func (this *TUISuite) TestRenderPlainTextBoard() {
	this.So(PlainText(new(Position)), should.Equal, Blank)
	this.So(PlainText(StartingPosition()), should.Equal, Starting)
	this.So(PlainText(must.Value(ParseFEN(RuyLopezFEN))), should.Equal, RuyLopez)
	this.So(PlainText(must.Value(ParseFEN(MiddleGameFEN))), should.Equal, MiddleGame)
}
