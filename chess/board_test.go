package chess

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestBitBoardSuite(t *testing.T) {
	should.Run(&BitBoardSuite{T: should.New(t)}, should.Options.UnitTests())
}

type BitBoardSuite struct {
	*should.T
}

func (this *BitBoardSuite) Test() {
	board := BitBoard(0)
	this.So(board.IsOccupied(1), should.BeFalse)
	this.Println(board)

	board.Occupy(1)
	this.So(board.IsOccupied(1), should.BeTrue)
	this.Println(board)
}
