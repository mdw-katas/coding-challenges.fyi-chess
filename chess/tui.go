package chess

import "strings"

func RenderPlainTextBoard(this *Position) string {
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
		if (s+1)%boardWidth == 0 {
			result.WriteString("\n")
		} else {
			result.WriteString(" ")
		}
	}
	return strings.TrimSpace(result.String())
}
