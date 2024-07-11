package chess

type Move struct {
	Color     Color
	Piece     Piece
	Source    Square
	Target    Square
	Promotion Piece
}
