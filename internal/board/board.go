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

// CreateStartpos returns an array of 120 squares, representing a 1D-array for the board state
// The game board is sourrounded by squares wich are invalid. This is needed for move generation to chech,
// if a piece moves out of bounds
//
// The invalid squares are the lowest ang highest two rows [0]-[19] & [100]-[119]
// as well as the left and rightmost files.
func CreateStartpos() []Square {
	board := make([]Square, 120)
	// Setting the valid squares; invalid ones are automatically invalid (zero value of bool is false)
	for i := 21; i < 98; i++ {
		if i%10 != 0 && i%10 != 9 {
			board[i].IsValid = true
		}
	}
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
