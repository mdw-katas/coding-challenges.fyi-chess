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
	this.So(RenderPlainTextBoard(new(Position)), should.Equal, Blank)
	this.So(RenderPlainTextBoard(StartingPosition()), should.Equal, Starting)
	this.So(RenderPlainTextBoard(must.Value(ParseFEN(RuyLopezFEN))), should.Equal, RuyLopez)
	this.So(RenderPlainTextBoard(must.Value(ParseFEN(MiddleGameFEN))), should.Equal, MiddleGame)
}
