package main

import (
	"fmt"
	"github.com/style84/cognichess/internal/board"
)

func main() {
	fmt.Println("Corgnichess is a minimal chess engine written in GO by Ingo Krause.")

	var positionBoard []board.Square = board.CreateStartpos()

	for i, sq := range positionBoard {
		fmt.Println("Square", i, "is valid =", sq.IsValid)
	}
}
