package main

import (
	"fmt"

	"mygolang/internal/chessdesk"
)

func main() {
	greeting()
	chessdesk.PrintChessDesk()
}

func greeting() {
	fmt.Println("Приветствую!")
}
