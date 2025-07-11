package board

type ColorType int
type PieceType int

const (
	White ColorType = iota
	Black
)

const (
	None PieceType = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	Kind     PieceType
	HasMoved bool
	Player   ColorType
}

type Square struct {
	OccupiedBy Piece
	EnPassant  bool
	IsValid    bool
}

func (s Square) IsEmpty() bool {
	if s.OccupiedBy.Kind == 0 {
		return true
	} else {
		return false
	}
}

func CreateStartpos() []Square {
	board := make([]Square, 120)

	// Setting up white Pawns
	for i := 21; i < 29; i++ {
		board[i].OccupiedBy.Kind = Pawn
		board[i].OccupiedBy.Player = White
	}
	// Setting up black Pawns
	for i := 81; i < 89; i++ {
		board[i].OccupiedBy.Kind = Pawn
		board[i].OccupiedBy.Player = Black
	}

	return board
}
