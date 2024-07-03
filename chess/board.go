package chess

type BitBoard int64

func (this BitBoard) IsOccupied(square int) bool {
	var mask uint64
	mask = 1 << uint64(this)
	return (uint64(square) & mask) == 0
}
func (this *BitBoard) Occupy(square int) {
	var mask BitBoard
	mask = BitBoard(1 << int64(square))
	*this |= mask
}
