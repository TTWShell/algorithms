package leetcode

import (
	"reflect"
	"testing"
)

func Test_findItinerary(t *testing.T) {
	tickets := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	if r := findItinerary(tickets); !reflect.DeepEqual(r, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}) {
		t.Fatal(r)
	}

	tickets = [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	if r := findItinerary(tickets); !reflect.DeepEqual(r, []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}) {
		t.Fatal(r)
	}
}
