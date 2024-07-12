package chess

type Color bool

const (
	White = Color(false)
	Black = Color(true)
)

func (this Color) Int() (result int) {
	if this == White {
		return 0
	}
	return 1
}
