package game

import (
	"fmt"
	"strings"
)

// Point - Describes a move in the game
type Point struct {
	X int
	Y int
}

// Data - Stores the data for the game
type Data struct {
	board         [][]string
	CurrentPlayer bool
	latestMove    Point
	turnCount     int
}

// InitGame - Set up a new game
func InitGame(startingPlayer bool) Data {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	var res = Data{
		board:         board,
		CurrentPlayer: startingPlayer,
		latestMove:    Point{0, 0},
		turnCount:     0,
	}
	return res
}

// PlayTurn - Make a move
func (game Data) PlayTurn(move Point) bool {
	x, y := move.X, move.Y

	// Out of Bounds
	if (x < 0 || x > 2) || (y < 0 || y > 2) {
		return false
	}

	// Square Taken
	if game.board[y][x] != "_" {
		return false
	}

	if game.CurrentPlayer {
		game.board[y][x] = "O"
	} else {
		game.board[y][x] = "X"
	}

	game.CurrentPlayer = !game.CurrentPlayer
	game.latestMove = move
	game.turnCount++
	return true
}

// CheckWin - Returns "_" if no winner, "O" if O wins, "X" if X wins and "tie" if it's a draw
func (game Data) CheckWin() string {
	refX, refY, player := game.latestMove.X, game.latestMove.Y, !game.CurrentPlayer

	var playerString string
	if player {
		playerString = "O"
	} else {
		playerString = "X"
	}

	// Check verticals
	for y := 0; y < 3; y++ {
		if game.board[y][refX] != playerString {
			break
		} else if y == 2 {
			return playerString
		}
	}

	// Check horizontals
	for x := 0; x < 3; x++ {
		if game.board[refY][x] != playerString {
			break
		} else if x == 2 {
			return playerString
		}
	}

	// Check diag 1
	if game.board[0][0] == playerString &&
		game.board[1][1] == playerString &&
		game.board[2][2] == playerString {
		return playerString
	}

	// Check diag 2
	if game.board[0][2] == playerString &&
		game.board[1][1] == playerString &&
		game.board[2][0] == playerString {
		return playerString
	}

	// Tie
	if game.turnCount == 9 {
		return "tie"
	}

	// No winner or tie
	return "_"
}

func (game Data) printBoard() {
	// / 0 1 2
	// 0 _ _ _
	// 1 _ _ _
	// 2 _ _ _

	fmt.Println("/ 0 1 2")
	for y := 0; y < 3; y++ {
		fmt.Printf("%d %s\n", y, strings.Join(game.board[y], " "))
	}
	fmt.Println()
}
