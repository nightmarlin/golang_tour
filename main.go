package main

import (
	"bufio"
	"fmt"
	game "github.com/Nightmarlin/golang_tour/x_and_o"
	"os"
	"strconv"
)

func main() {
	// fmt.Println(greeterTime())
	// printRandomInt(22)
	// printSqrt(34)
	// fmt.Printf("1 + 2 = %d\go n2 - 1 = %d\n", add(1, 2), sub(2, 1))
	// fmt.Printf("The 17th triangular number is %d\n", sum(0, 17))
	// fmt.Printf("The 3rd and 14th fibonnacci numbers are %d and %d\n", fib(3), fib(14))
	// fmt.Printf("The sqrt of 234 is %g after 5 iterations and %g after 15", sqrt(234, 5), sqrt(234, 15))
	// deference()
	// testvar := 23434254643245
	// pointerisms(testvar)
	// tim := Person{"Tim", 34, Address{23, "Artemy Road", "LN1 0XX", "London", "England"}}
	// fmt.Println(tim)

	gameData := game.InitGame(true)

	fmt.Println(gameData)

	for (gameData.CheckWin()) == "_" {
		var currentPlayerString = ""
		if gameData.CurrentPlayer {
			currentPlayerString = "O"
		} else {
			currentPlayerString = "X"
		}

		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("It's %s's turn!\n X ==> ", currentPlayerString)
		resX, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("\n Y ==> ")
		resY, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
			return
		}

		xValid, x := validateInput(string(resX))
		yValid, y := validateInput(string(resY))

		if !xValid || yValid {
			fmt.Println("Error: Invalid argument")
			return
		}

		gameData.PlayTurn(game.Point{X: x, Y: y})
	}
}

func validateInput(in string) (bool, int) {
	switch in {
	case "0":
		fallthrough
	case "1":
		fallthrough
	case "2":
		res, err := strconv.Atoi(in)
		if err != nil {
			return false, 0
		}
		return true, res

	default:
		return false, 0
	}
}
