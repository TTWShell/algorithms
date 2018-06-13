package lstring

import "testing"

func Test_checkRecord(t *testing.T) {
	input := []string{"LALL", "PPALLP", "PPALLL"}
	output := []bool{true, true, false}

	for i := range input {
		if r := checkRecord(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
