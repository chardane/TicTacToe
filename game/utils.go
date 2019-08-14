package game

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseMoveInput(move string) (row int, col int, err error) {
	var rowErr, colErr error
	move = strings.TrimSpace(move)
	moves := strings.Split(move, ",")

	if len(moves) != 2 {
		return 0, 0, fmt.Errorf("Invalid move.")
	}

	row, rowErr = strconv.Atoi(moves[0])
	if rowErr != nil {
		return 0, 0, fmt.Errorf("Invalid row.")
	}

	col, colErr = strconv.Atoi(moves[1])
	if colErr != nil {
		return 0, 0, fmt.Errorf("Invalid col.")
	}

	return row, col, nil
}

func IsPlayerX(turnCount int) bool {
	if (turnCount % 2) == 0 {
		return true
	}
	return false
}

func printRow(rowLabel string, row []string) {
	labelSpacer := "  "
	colSpacer := " | "
	displayRow := rowLabel + labelSpacer
	for i := range row {
		display := row[i]
		if isNotLastIndex(i, len(row)) {
			display += colSpacer
		}

		displayRow += display
	}
	fmt.Println(displayRow)
}

func printLabelRow(size int) {
	colSpacer := "   "
	displayRow := colSpacer

	for i := 0; i < size; i++ {
		displayRow += strconv.Itoa(i)
		if isNotLastIndex(i, size) {
			displayRow += colSpacer
		}
	}
	fmt.Printf("\n%s\n", displayRow)
}

func printSpacerRow(size int) {
	colSpacer := "---"
	displayRow := "   "

	for i := 0; i < size; i++ {
		displayRow += "-"
		if isNotLastIndex(i, size) {
			displayRow += colSpacer
		}
	}
	fmt.Println(displayRow)
}

func isNotLastIndex(index int, length int) bool {
	if index < length-1 {
		return true
	}
	return false
}

