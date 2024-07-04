package chess

import (
	"fmt"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestBitBoard(t *testing.T) {
	board := BitBoard(0)
	should.So(t, board, should.Equal, 0)
	for square := 0; square < 64; square++ {
		t.Run(fmt.Sprint(square), func(t *testing.T) {
			should.So(t, board.IsOccupied(square), should.BeFalse)
			board.Occupy(square)
			should.So(t, board.IsOccupied(square), should.BeTrue)
		})
	}
	should.So(t, board, should.Equal, BitBoard(0xffffffffffffffff))
}
