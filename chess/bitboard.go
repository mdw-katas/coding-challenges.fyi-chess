package chess

import "math/bits"

type BitBoard uint64

func NewBitBoard(squares ...Square) *BitBoard {
	result := new(BitBoard)
	for _, square := range squares {
		result.Occupy(square)
	}
	return result
}
func (this *BitBoard) mask(bit Square) BitBoard {
	return BitBoard(1 << bit)
}
func (this *BitBoard) IsOccupied(square Square) bool {
	return *this&this.mask(square) > 0
}
func (this *BitBoard) Occupy(square Square) {
	*this |= this.mask(square)
}
func (this *BitBoard) Vacate(square Square) {
	*this &= ^this.mask(square)
}
func (this *BitBoard) Flip() {
	if *this > 0 {
		*this = BitBoard(bits.Reverse64(uint64(*this)))
	}
}
func (this *BitBoard) OccupiedSquares() func(func(Square) bool) {
	return func(yield func(Square) bool) {
		for _, square := range allSquares { // TODO: optimize to avoid checking every square...
			if this.IsOccupied(square) {
				yield(square)
			}
		}
	}
}
