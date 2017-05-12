package leetcode

import "testing"

func Test_addBinary(t *testing.T) {
	as := []string{"0", "0", "1", "11", "11", "1111"}
	bs := []string{"0", "1", "1", "10", "1", "1111"}
	rs := []string{"0", "1", "10", "101", "100", "11110"}
	for i := 0; i < len(as); i++ {
		if r := addBinary(as[i], bs[i]); r != rs[i] {
			t.Log(as[i], bs[i], rs[i], r)
			t.Fail()
		}
	}
}
