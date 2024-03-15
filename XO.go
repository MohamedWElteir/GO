package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var board [3][3]int32
var currentPlayer  = 'X'
var row, col int

func main() {
	initializeBoard()
	displayBoard()

	for {
		getPlayerMove(row, col)
		checkWinner(currentPlayer)
		currentPlayer = switchPlayer(currentPlayer)
	}

}

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = '_'
		}
	}
}

func displayBoard() {
	clearScreen()
	for i := 0; i < 3; i++ {
		row := make([]string, 3) 
		for j := 0; j < 3; j++ {
			row[j] = string(board[i][j]) 
		}
		fmt.Println(strings.Join(row, " "))
	}
	fmt.Println()
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getPlayerMove(row, col int) {
	fmt.Printf("Player %c - Enter your move (row column): ", currentPlayer)
	fmt.Scan(&row, &col)
	if isValidMove(row-1, col-1) {
		makeMove(row-1, col-1, currentPlayer)
		displayBoard()
	} else {
		fmt.Println("Invalid move. Please try again.")
		getPlayerMove(row, col)
	}
}

func isValidMove(row, col int) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 || board[row][col] != '_' {
		return false
	}
	return true
}

func makeMove(row, col int, player int32) {
	board[row][col] = player
}

func isWinner(player int32) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

func isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == '_' {
				return false
			}
		}
	}
	return true
}

func switchPlayer(currentPlayer int32) int32 {
	if currentPlayer == 'X' {
		return 'O'
	} else {
		return 'X'
	}
}

func checkWinner(player int32) {
	if isWinner(player) {
		fmt.Printf("Player %c wins!\n", player)
		os.Exit(0)
	} else if isBoardFull() {
		fmt.Println("It's a draw!")
		os.Exit(0)
	}
}