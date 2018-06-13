package stack

import "testing"

func Test_IsBalanceSymbol(t *testing.T) {
	if r := IsBalanceSymbol("[()]"); r != true {
		t.Fatal(r)
	}

	if r := IsBalanceSymbol("[(]}"); r != false {
		t.Fatal(r)
	}

	if r := IsBalanceSymbol("]"); r != false {
		t.Fatal(r)
	}

	if r := IsBalanceSymbol("[({])}}"); r != false {
		t.Fatal(r)
	}

	if r := IsBalanceSymbol("(a+b){[d]c*d}{"); r != false {
		t.Fatal(r)
	}
}
