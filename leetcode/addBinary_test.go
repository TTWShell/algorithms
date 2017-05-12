package leetcode

import "testing"

func Test_addBinary(t *testing.T) {
	as := []string{"0", "1", "11", "11"}
	bs := []string{"1", "1", "10", "1"}
	rs := []string{"1", "10", "101", "100"}
	for i := 0; i < len(as); i++ {
		if r := addBinary(as[i], bs[i]); r != rs[i] {
			t.Log(as[i], bs[i], rs[i], r)
			t.Fail()
		}
	}
}
