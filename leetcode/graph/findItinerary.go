/* https://leetcode.com/problems/reconstruct-itinerary/description/
Given a list of airline tickets represented by pairs of departure and arrival airports [from, to], reconstruct the itinerary in order.
All of the tickets belong to a man who departs from JFK. Thus, the itinerary must begin with JFK.

Note:
    1. If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order when read as a single string.
        For example, the itinerary ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"].
    2. All airports are represented by three capital letters (IATA code).
    3. You may assume all tickets form at least one valid itinerary.

Example 1:
    tickets = [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
    Return ["JFK", "MUC", "LHR", "SFO", "SJC"].

Example 2:
    tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
    Return ["JFK","ATL","JFK","SFO","ATL","SFO"].
    Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"]. But it is larger in lexical order.
*/

package leetcode

import (
	"sort"
	"strings"
)

func findItinerary(tickets [][]string) []string {
	notIn := func(list []int, target int) bool {
		for _, num := range list {
			if num == target {
				return false
			}
		}
		return true
	}

	type ticketIndex map[string]int
	graph := make(map[string]ticketIndex, len(tickets))
	for index, ticket := range tickets {
		if _, ok := graph[ticket[0]]; !ok {
			graph[ticket[0]] = make(ticketIndex)
		}
		graph[ticket[0]][ticket[1]] = index
	}

	var helper func(graph map[string]ticketIndex, used []int, now string) [][]string
	helper = func(graph map[string]ticketIndex, used []int, now string) [][]string {
		if len(used) == len(tickets) {
			return [][]string{{now}}
		}
		res := [][]string{}
		for to, index := range graph[now] {
			if !notIn(used, index) {
				continue
			}
			result := helper(graph, append(used, index), to)
			if len(result) != 0 {
				for _, sub := range result {
					res = append(res, append([]string{now}, sub...))
				}
			}
		}
		return res
	}

	allRes := helper(graph, []int{}, "JFK")
	tmp := make([]string, len(allRes), len(allRes))
	for i := range allRes {
		tmp[i] = strings.Join(allRes[i], "")
	}
	sort.Strings(tmp)

	res := make([]string, len(tickets)+1, len(tickets)+1)
	for i := 0; i < len(tmp[0]); i += 3 {
		res[i/3] = string(tmp[0][i : i+3])
	}
	return res
}
