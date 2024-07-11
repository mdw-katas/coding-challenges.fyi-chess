package chess

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestBitBoard(t *testing.T) {
	board := BitBoard(0)
	should.So(t, board, should.Equal, 0)
	for _, square := range allSquares {
		t.Run("occupy-"+square.String(), func(t *testing.T) {
			should.So(t, board.IsOccupied(square), should.BeFalse)
			board.Occupy(square)
			should.So(t, board.IsOccupied(square), should.BeTrue)
		})
	}
	should.So(t, board, should.Equal, BitBoard(0xffffffffffffffff))
	t.Log(board)
	for _, square := range allSquares {
		t.Run("vacate-"+square.String(), func(t *testing.T) {
			t.Log(board)
			should.So(t, board.IsOccupied(square), should.BeTrue)
			board.Vacate(square)
			should.So(t, board.IsOccupied(square), should.BeFalse)
		})
	}
}
