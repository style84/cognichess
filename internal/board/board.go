package board

type Color int
type PieceType int

const (
	White Color = iota
	Black
)

const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	Kind     PieceType
	HasMoved bool
	Player   Color
}
