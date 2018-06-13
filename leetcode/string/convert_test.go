package lstring

import "testing"

func Test_convert(t *testing.T) {
	if r := convert("PAYPALISHIRING", 3); r != "PAHNAPLSIIGYIR" {
		t.Fatal(r)
	}
}
