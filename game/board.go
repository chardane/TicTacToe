package game

import (
	"fmt"
	"math"
	"strconv"
)

// Game pieces
const (
	X = "x"
	O = "o"
  Empty = " "
)

type GameResult int
const(
	PlayersTie GameResult = 0
	PlayerXWins GameResult = 1
	PlayerOWins GameResult = 2
)

type Board struct {
	Rows   [][]string
	Winner *GameResult
	size   int //can only be a square board of any size
}

// function types used for finding wins in rows or columns
type FirstPieceFinder func([][]string, int,int) string
type CurrentPieceFinder func([][]string, int,int) string

// function type used for finding wins in diagonals (forward and backward)
type RowIterator func(int) int

func (b *Board) CreateEmptyBoard(size int){
	b.size = size
	rows := make([][]string, size)

	for i := range rows {
		rows[i] = make([]string, size)
		for j := range rows[i] {
			rows[i][j] = Empty
		}
	}

	b.Rows = rows
}

func (b *Board) PrintBoard() {
	printLabelRow(b.size)

	for i, row := range b.Rows {
		var rowLabel string
		rowLabel = strconv.Itoa(i)

		printRow(rowLabel, row)
		if isNotLastIndex(i, len(b.Rows)){
			printSpacerRow(b.size)
		}
	}
}

func (b *Board) PlaceMoveAndCheckWin(playerX bool, row int, col int, turnCount int) error {
	piece := O
	if playerX {
		piece = X
	}
	err := b.placeMove(piece, row, col)
	if err != nil {
		return err
	}

	winner := b.findWinner(piece, row, col)
	if winner != nil {
		b.setWinner(*winner)
		return nil
	}

	// if no winner, check if board is full
	if b.isBoardFull(turnCount) {
		draw := PlayersTie
		b.Winner = &draw
		return nil
	}

	// if no winner and board is not full
	return nil
}

func (b *Board) CongratulateWinner() {
	if b.Winner == nil {
		fmt.Println("It's a tie!")
		return
	}
	switch *b.Winner {
	case PlayerXWins:
		fmt.Println("Congratulations, Player X!")
	case PlayerOWins:
		fmt.Println("Congratulations, Player O!")
	case PlayersTie:
		fmt.Println("It's a tie!")
	}
}

func (b *Board) findWinner(inputPiece string, row int, col int) *string {
	// check rows
	rowFirstPieceFinder := func(boardRows [][]string, outerIndex int, innerIndex int) string {
		return boardRows[outerIndex][0]
	}
	rowPieceFinder := func(boardRows [][]string, outerIndex int, innerIndex int) string {
		return boardRows[outerIndex][innerIndex]
	}
	winner := b.checkRowOrColsForWins(rowFirstPieceFinder, rowPieceFinder)
	if winner != nil {
		return winner
	}

	// check cols
	colFirstPieceFinder := func(boardRows [][]string, outerIndex int, innerIndex int) string {
		return boardRows[0][outerIndex]
	}
	colPieceFinder := func(boardRows [][]string, outerIndex int, innerIndex int) string {
		return boardRows[innerIndex][outerIndex]
	}
	winner = b.checkRowOrColsForWins(colFirstPieceFinder, colPieceFinder)
	//winner = b.checkColsForWins()
	if winner != nil {
		return winner
	}

	// check forward diagonal (/)
	forwardDiagonalIterator := func(row int) int {
		return row+1
	}
	winner = b.checkDiagonalForWins(0, forwardDiagonalIterator)
	if winner != nil {
		return winner
	}

	// check backward diagonal (\)
	backwardDiagonalIterator := func(row int) int {
		return row-1
	}
	winner = b.checkDiagonalForWins(b.size-1, backwardDiagonalIterator)
	if winner != nil {
		return winner
	}

	return nil
}

func (b *Board) checkDiagonalForWins(initialRowIndex int, rowIterator RowIterator) *string {
	firstPiece := b.Rows[initialRowIndex][0]
	if firstPiece == Empty {
		return nil // if the firstPiece is Empty, this diagonal is not complete
	}
	rowIndex := initialRowIndex
	for colIndex:=0; colIndex < b.size; colIndex++ {
		piece := b.Rows[rowIndex][colIndex]
		if piece != firstPiece {
			return nil
		}
		rowIndex = rowIterator(rowIndex)
	}
	return &firstPiece
}

func (b *Board) checkRowOrColsForWins(firstPieceFinder FirstPieceFinder, currentPieceFinder CurrentPieceFinder) *string {
	matchingPieces := 0
	innerIndex := 0
	for outerIndex:=0; outerIndex < b.size; outerIndex++ {
		firstPiece := firstPieceFinder(b.Rows, outerIndex, innerIndex)
		matchingPieces = 0
		if firstPiece == Empty {
			continue // if the firstPiece is Empty, this col is not complete, so move on to next col
		}
		for innerIndex = 0; innerIndex < b.size; innerIndex++ {
			// check that each piece is the same as the first
			// which is sufficient to know that the entire row is the same piece
			// note that matchingPieces will always increment for the first element
			piece := currentPieceFinder(b.Rows, outerIndex, innerIndex)
			if piece == firstPiece {
				matchingPieces++
				if matchingPieces == b.size {
					return &piece
				}
			} else {
				break // move onto the next column
			}
		}
	}
	return nil
}

func (b *Board) isBoardFull(turnCount int) bool {
	totalAvailableMoves := math.Pow(float64(b.size), float64(2))
	// since we increment turnCount outside of the Board
	// check if we currently just used our last possible move
	if turnCount == int(totalAvailableMoves)-1 {
		return true
	}
	return false
}

func (b *Board) placeMove(piece string, row int, col int) error {
	if b.Rows[row][col] != Empty {
		return fmt.Errorf("Invalid move, space not empty.")
	}
	b.Rows[row][col] = piece
	return nil
}

func (b *Board) setWinner(winner string){
	switch winner {
	case X:
		won := PlayerXWins
		b.Winner = &won
	case O:
		won := PlayerOWins
		b.Winner = &won
	}
}
