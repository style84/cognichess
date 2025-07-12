package main

import (
	"fmt"
	"github.com/style84/cognichess/internal/board"
)

func main() {
	fmt.Println("Corgnichess is a minimal chess engine written in GO by Ingo Krause.")

	var mailbox []board.Square = board.CreateStartpos()

	for i, sq := range mailbox {
		fmt.Println("Square", i, "is valid =", sq.IsValid)
	}
}
