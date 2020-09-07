package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("Happy New Year! \n")

	createChristmasTree(5)

	fmt.Printf("Tic-Tac-Toe Board: \n")

	createTicTacToeBoard()

}

func createChristmasTree(line uint8) {
	var i, j, k uint8

	// fmt.Print("Enter line number: ")

	// fmt.Scanln(&line)

	for i = 1; i <= line; i++ {
		fmt.Println()
		for j = 1; j <= (line - i); j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}

	}
	fmt.Println()
}

func createTicTacToeBoard() {

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// fill board
	board[0][0] = "   |"
	board[0][1] = "   "
	board[0][2] = "|   "
	board[1][0] = "---+"
	board[1][1] = "---"
	board[1][2] = "+---"
	board[2][0] = "   |"
	board[2][1] = "   "
	board[2][2] = "|   "
	board[3][0] = "---+"
	board[3][1] = "---"
	board[3][2] = "+---"
	board[4][0] = "   |"
	board[4][1] = "   "
	board[4][2] = "|   "

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], ""))
	}
}
