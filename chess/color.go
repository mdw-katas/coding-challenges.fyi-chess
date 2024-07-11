package chess

type Color bool

const (
	White = Color(true)
	Black = Color(false)
)

func (this Color) Int() (result int) {
	if this == White {
		return 0
	}
	return 1
}
