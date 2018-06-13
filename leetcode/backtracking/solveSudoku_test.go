package lbacktracking

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_solveSudoku(t *testing.T) {
	initBoard := func(input []string) [][]byte {
		board := [][]byte{}
		for i := range input {
			board = append(board, []byte(input[i]))
		}
		return board
	}

	printBoard := func(board [][]byte) {
		for r := range board {
			for c := range board[r] {
				fmt.Printf("%c ", board[r][c])
			}
			fmt.Println()
		}
	}

	board := initBoard([]string{"..9748...", "7........", ".2.1.9...", "..7...24.", ".64.1.59.", ".98...3..", "...8.3.2.", "........6", "...2759.."})
	output := initBoard([]string{"519748632", "783652419", "426139875", "357986241", "264317598", "198524367", "975863124", "832491756", "641275983"})
	if solveSudoku(board); !reflect.DeepEqual(board, output) {
		printBoard(board)
		t.Fail()
	}
}
