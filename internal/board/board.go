package board

import (
	"errors"
	"strconv"
	"strings"
)

type ColorType int
type PieceType int
type fenGroup int

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

const (
	position fenGroup = iota
	activeColor
	castlingAvailability
	enPassant
	halfMoveClk
	moveNumber
)

type Piece struct {
	Kind     PieceType
	HasMoved bool
	Player   ColorType
}

type Square struct {
	OccupiedBy  Piece
	IsEnPassant bool
	IsValid     bool
}

func (s Square) IsEmpty() bool {
	if s.OccupiedBy.Kind == 0 {
		return true
	} else {
		return false
	}
}

// Rows is a helper map to iterate over valid squares in each row on the chessboard
//
// key: int [1-8] are the rows as marked on the board
// values: []int are slices containing the idexes of the corresponding squares in the board slice
var rows = map[int][]int{
	8: {91, 92, 93, 94, 95, 96, 97, 98},
	7: {81, 82, 83, 84, 85, 86, 87, 88},
	6: {71, 72, 73, 74, 75, 76, 77, 78},
	5: {61, 62, 63, 64, 65, 66, 67, 68},
	4: {51, 52, 53, 54, 55, 56, 57, 58},
	3: {41, 42, 43, 44, 45, 46, 47, 48},
	2: {31, 32, 33, 34, 35, 36, 37, 38},
	1: {21, 22, 23, 24, 25, 26, 27, 28},
}

func CreateFromFEN(fen string) ([]Square, error) {
	board := make([]Square, 120)
	// Checking if the FEN is valid
	fenGroups := strings.Fields(fen)
	if len(fenGroups) != 6 {
		return board, errors.New("FEN token number mismatch")
	}

	positionStrings := strings.Split(fenGroups[position], "/")
	if len(positionStrings) != 8 {
		return board, errors.New("FEN row number mismatch.")
	}

	// Setting the valid squares; invalid ones are automatically invalid (zero value of bool is false)
	for i := 21; i < 98; i++ {
		if i%10 != 0 && i%10 != 9 {
			board[i].IsValid = true
		}
	}

	// Going through all rows starting from the last row [8]
	activeRow := 8
	for j := range 8 {
		emptySquares := 0
		for i, s := range positionStrings[j] {
			if emptySquares > 0 {
				emptySquares--
				continue
			}
			switch string(s) {
			case "k":
				board[rows[activeRow][i]].OccupiedBy.Kind = King
				board[rows[activeRow][i]].OccupiedBy.Player = Black

			case "q":
				board[rows[activeRow][i]].OccupiedBy.Kind = Queen
				board[rows[activeRow][i]].OccupiedBy.Player = Black

			case "r":
				board[rows[activeRow][i]].OccupiedBy.Kind = Rook
				board[rows[activeRow][i]].OccupiedBy.Player = Black

			case "b":
				board[rows[activeRow][i]].OccupiedBy.Kind = Bishop
				board[rows[activeRow][i]].OccupiedBy.Player = Black

			case "n":
				board[rows[activeRow][i]].OccupiedBy.Kind = Knight
				board[rows[activeRow][i]].OccupiedBy.Player = Black

			case "p":
				board[rows[activeRow][i]].OccupiedBy.Kind = Pawn
				board[rows[activeRow][i]].OccupiedBy.Player = Black

			case "K":
				board[rows[activeRow][i]].OccupiedBy.Kind = King
				board[rows[activeRow][i]].OccupiedBy.Player = White

			case "Q":
				board[rows[activeRow][i]].OccupiedBy.Kind = Queen
				board[rows[activeRow][i]].OccupiedBy.Player = White

			case "R":
				board[rows[activeRow][i]].OccupiedBy.Kind = Rook
				board[rows[activeRow][i]].OccupiedBy.Player = White

			case "B":
				board[rows[activeRow][i]].OccupiedBy.Kind = Bishop
				board[rows[activeRow][i]].OccupiedBy.Player = White

			case "N":
				board[rows[activeRow][i]].OccupiedBy.Kind = Knight
				board[rows[activeRow][i]].OccupiedBy.Player = White

			case "P":
				board[rows[activeRow][i]].OccupiedBy.Kind = Pawn
				board[rows[activeRow][i]].OccupiedBy.Player = White

			default:
				// s is not a piece, must be number of squares to skip
				toSkip, _ := strconv.Atoi(string(s))
				emptySquares += toSkip
			}
		}
		activeRow--
	}

	return board, nil
}

// CreateStartpos returns an array of 120 squares, representing a 1D-array for the board state
// The game board is sourrounded by squares wich are invalid. This is needed for move generation to chech,
// if a piece moves out of bounds
//
// The invalid squares are the lowest ang highest two rows [0]-[19] & [100]-[119]
// as well as the left and rightmost files.
func CreateStartpos() []Square {
	startPosFEN := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	board := make([]Square, 120)
	board, _ = CreateFromFEN(startPosFEN)
	return board
}
