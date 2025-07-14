package board

import (
	//	"fmt"
	"testing"
)

func TestCreateFromFEN(t *testing.T) {
	correctFEN := ("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	wrongTokenFEN := ("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR KQkq - 0 1")
	wrongPositionFEN := ("pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	b := make([]Square, 120)
	b, err := CreateFromFEN(correctFEN)
	if err != nil {
		t.Errorf("Call with correct FEN; want nil, got %s", err)
	}

	var checkedPiece PieceType
	checkedPiece = b[21].OccupiedBy.Kind
	if checkedPiece != Rook {
		t.Errorf("Check for correct square occupation. Wanted %d, got %d", Rook, checkedPiece)
	}
	checkedPiece = b[64].OccupiedBy.Kind
	if checkedPiece != None {
		t.Errorf("Check for correct square occupation. Wanted %d, got %d", None, checkedPiece)
	}
	checkedPiece = b[31].OccupiedBy.Kind
	if checkedPiece != Pawn {
		t.Errorf("Check for correct square occupation. Wanted %d, got %d", Pawn, checkedPiece)
	}

	_, err = CreateFromFEN(wrongTokenFEN)
	if err == nil {
		t.Errorf("Call with incorrect FEN; want !nil, got %s", err)
	}

	_, err = CreateFromFEN(wrongPositionFEN)
	if err == nil {
		t.Errorf("Call with incorrect FEN; want !nil, got %s", err)
	}
}
