package chess

import (
	"errors"
	"fmt"
	"strings"
)

type Position struct {
	WhiteToMove   bool
	White         [6]BitBoard
	Black         [6]BitBoard
	Castling      BitBoard
	EnPassant     BitBoard
	HalfMoveClock int
	FullMoveCount int
}

func StartingPosition() *Position {
	result, err := ParseFEN(startingFEN)
	if err != nil {
		panic(err) // TODO: remove when finished with ParseFEN (this should never happen)
	}
	return result
}

func ParseFEN(fen string) (result *Position, err error) {
	result = new(Position)
	fields := strings.Split(fen, " ")
	if len(fields) != 6 {
		return nil, fmt.Errorf("%w: want 6 fields, had %d instead", errInvalidFEN, len(fields))
	}
	ranks := strings.Split(fields[0], "/")
	if len(ranks) != 8 {
		return nil, fmt.Errorf("%w: want 8 ranks, had %d instead", errInvalidFEN, len(ranks))
	}
	for r, rank := range ranks {
		r = 7 - r
		offset := 0
		for c, char := range rank {
			square := Square((8 * r) + c + offset)
			switch char {
			case '1', '2', '3', '4', '5', '6', '7', '8':
				offset += int(char - '0')
			case 'K', 'Q', 'R', 'B', 'N', 'P':
				result.White[fen2type[char]].Occupy(square)
			case 'k', 'q', 'r', 'b', 'n', 'p':
				result.Black[fen2type[char]].Occupy(square)
			default:
				return nil, fmt.Errorf("%w: invalid character in piece placement section '%c'", errInvalidFEN, char)
			}
		}
	}
	switch playerToMove := fields[1]; playerToMove {
	case "w":
		result.WhiteToMove = true
	case "b":
		result.WhiteToMove = false
	default:
		return nil, fmt.Errorf("%w: invalid value in player-to-move section '%s'", errInvalidFEN, fields[1])
	}
	castling := fields[2]
	if len(castling) > 0 && castling[0] == 'K' {
		result.Castling.Occupy(whiteKingsideCastleTarget)
		castling = castling[1:]
	}
	if len(castling) > 0 && castling[0] == 'Q' {
		result.Castling.Occupy(whiteQueensideCastleTarget)
		castling = castling[1:]
	}
	if len(castling) > 0 && castling[0] == 'k' {
		result.Castling.Occupy(blackKingsideCastleTarget)
		castling = castling[1:]
	}
	if len(castling) > 0 && castling[0] == 'q' {
		result.Castling.Occupy(blackQueensideCastleTarget)
		castling = castling[1:]
	}
	if len(castling) > 0 && castling != "-" {
		return nil, fmt.Errorf("%w: invalid value in castling section '%s'", errInvalidFEN, fields[2])
	}
	if enPassant := fields[3]; enPassant != "-" {
		parseEnPassant, _ := parseSquare(enPassant)
		switch parseEnPassant {
		case a3, b3, c3, d3, e3, f3, g3, h3, a6, b6, c6, d6, e6, f6, g6, h6:
		default:
			return nil, fmt.Errorf("%w: en passant target '%s'", errInvalidFEN, enPassant)
		}
		result.EnPassant.Occupy(parseEnPassant)
	}
	return result, err
}

var errInvalidFEN = errors.New("invalid FEN")

func (this *Position) String() string {
	var result strings.Builder
	for s, square := range allSquares {
		length := result.Len()
		for _, pieceType := range allPieceTypes {
			if this.White[pieceType].IsOccupied(square) {
				result.WriteString(pieceType.WhiteFigurine())
			} else if this.Black[pieceType].IsOccupied(square) {
				result.WriteString(pieceType.BlackFigurine())
			}
		}
		if result.Len() == length {
			result.WriteString("-")
		}
		if (s+1)%8 == 0 {
			result.WriteString("\n")
		} else {
			result.WriteString(" ")
		}
	}
	return strings.TrimSpace(result.String())
}

const startingFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
