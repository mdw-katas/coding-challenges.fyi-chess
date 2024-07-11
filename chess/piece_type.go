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

var initial2type = map[rune]PieceType{
	'K': King /****/, 'k': King,
	'Q': Queen /***/, 'q': Queen,
	'R': Rook /****/, 'r': Rook,
	'B': Bishop /**/, 'b': Bishop,
	'N': Knight /**/, 'n': Knight,
	'P': Pawn /****/, 'p': Pawn,
}

var allPieceTypes = []PieceType{King, Queen, Rook, Bishop, Knight, Pawn}

func (this PieceType) WhiteInitial() string {
	switch this {
	case King:
		return "K"
	case Queen:
		return "Q"
	case Rook:
		return "R"
	case Bishop:
		return "B"
	case Knight:
		return "N"
	case Pawn:
		return "P"
	default:
		return "?"
	}
}
func (this PieceType) BlackInitial() string {
	switch this {
	case King:
		return "k"
	case Queen:
		return "q"
	case Rook:
		return "r"
	case Bishop:
		return "b"
	case Knight:
		return "n"
	case Pawn:
		return "p"
	default:
		return "¿"
	}
}
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
		return "♟"
	default:
		return "¿"
	}
}
