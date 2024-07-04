package chess

type BitBoard uint64

func (this *BitBoard) mask(bit Square) BitBoard {
	return BitBoard(1 << bit)
}
func (this *BitBoard) IsOccupied(square Square) bool {
	return *this&this.mask(square) > 0
}
func (this *BitBoard) Occupy(square Square) {
	*this |= this.mask(square)
}
