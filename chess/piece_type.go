package chess

type PieceType int

const (
	King PieceType = iota
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

func (this PieceType) WhiteFigurine() string {
	switch this {
	case King:
		return "♔"
	case Queen:
		return "♕"
	case Rook:
		return "♖"
	case Bishop:
		return "♗"
	case Knight:
		return "♘"
	case Pawn:
		return "♙"
	default:
		return "?"
	}
}
func (this PieceType) BlackFigurine() string {
	switch this {
	case King:
		return "♚"
	case Queen:
		return "♛"
	case Rook:
		return "♜"
	case Bishop:
		return "♝"
	case Knight:
		return "♞"
	case Pawn:
		return "♟︎"
	default:
		return "¿"
	}
}
