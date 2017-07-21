package leetcode

import "testing"

func Test_isValidSudoku(t *testing.T) {
	board := [][]byte{}
	input := []string{".87654321", "2........", "3........", "4........", "5........", "6........", "7........", "8........", "9........"}
	for i := range input {
		board = append(board, []byte(input[i]))
	}
	if r := isValidSudoku(board); r != true {
		t.Fatal(r)
	}
}
