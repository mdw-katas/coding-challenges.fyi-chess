package chess

type BitBoard uint64

func (this *BitBoard) mask(bit int) BitBoard {
	return BitBoard(1 << uint64(bit))
}
func (this *BitBoard) IsOccupied(square int) bool {
	return *this&this.mask(square) > 0
}
func (this *BitBoard) Occupy(square int) {
	*this |= this.mask(square)
}
