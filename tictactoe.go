package main

import (
	"RC/tictactoe/game"
	"bufio"
	"fmt"
	"os"
)

func main() {
	printGameWelcomeBanner()
	turnCount := 0
	board := new(game.Board)
	board.CreateEmptyBoard(3)
	board.PrintBoard()

	for board.Winner == nil {
		playerX := game.IsPlayerX(turnCount)
		playerInput := collectInput(playerX)
		inputRow, inputCol, err := game.ParseMoveInput(playerInput)
		if err != nil {
			fmt.Println(err, "Please try again.")
			continue // let the current player try again
		}

		moveErr := board.PlaceMoveAndCheckWin(playerX, inputRow, inputCol, turnCount)
		if moveErr != nil {
			fmt.Println(moveErr, "Please try again.")
			continue // let the current player try again
		}

		board.PrintBoard()
		turnCount++
	}
	board.CongratulateWinner()
}

func collectInput(playerX bool) string {
	var sentence []byte
	var err error

	promptString := "Player O: "
	if playerX {
		promptString = "Player X: "
	}

	buf := bufio.NewReader(os.Stdin)
	fmt.Print("\n", promptString)

	sentence, err = buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}
	return string(sentence)
}

func printGameWelcomeBanner() {
	fmt.Println("*************************")
	fmt.Println("*                       *")
	fmt.Println("* Let's play TicTacToe! *")
	fmt.Println("*                       *")
	fmt.Println("*************************")
}
