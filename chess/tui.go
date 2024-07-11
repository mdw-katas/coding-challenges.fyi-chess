package chess

import "strings"

func RenderPlainTextBoard(this *Position) string {
	var result strings.Builder
	for s, square := range allSquares {
		length := result.Len()
		for _, color := range []Color{White, Black} {
			for _, pieceType := range allPieceTypes {
				if this.Pieces[color.Int()][pieceType].IsOccupied(square) {
					result.WriteString(pieceType.Figurine(color))
				}
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
