package lht

import "testing"

func Test_isValidSudoku(t *testing.T) {
	initBoard := func(input []string) [][]byte {
		board := [][]byte{}
		for i := range input {
			board = append(board, []byte(input[i]))
		}
		return board
	}
	board := initBoard([]string{"....5..1.", ".4.3.....", ".....3..1", "8......2.", "..2.7....", ".15......", ".....2...", ".2.9.....", "..4......"})
	if r := isValidSudoku(board); r != false {
		t.Fatal(r)
	}

	board = initBoard([]string{".87654321", "2........", "3........", "4........", "5........", "6........", "7........", "8........", "9........"})
	if r := isValidSudoku(board); r != true {
		t.Fatal(r)
	}

	board = initBoard([]string{"..5.....6", "....14...", ".........", ".....92..", "5....2...", ".......3.", "...54....", "3.....42.", "...27.6.."})
	if r := isValidSudoku(board); r != true {
		t.Fatal(r)
	}

	board = initBoard([]string{"..4...63.", ".........", "5......9.", "...56....", "4.3.....1", "...7.....", "...5.....", ".........", "........."})
	if r := isValidSudoku(board); r != false {
		t.Fatal(r)
	}
}
