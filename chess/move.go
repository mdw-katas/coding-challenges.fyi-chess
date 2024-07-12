package chess

type Move struct {
	Color     Color
	Piece     Piece
	Source    Square
	Target    Square
	Captured  Piece
	Promotion Piece
}
