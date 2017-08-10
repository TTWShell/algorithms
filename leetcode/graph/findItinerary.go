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

import "sort"

func findItinerary(tickets [][]string) []string {
	graph := make(map[string][]string, len(tickets))
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		if _, ok := graph[from]; !ok {
			graph[from] = []string{}
		}
		graph[from] = append(graph[from], to)
	}
	for _, ticket := range tickets {
		sort.Strings(graph[ticket[0]])
	}

	var helper func(now string, n int) (res []string)
	helper = func(now string, n int) (res []string) {
		if n == 0 {
			return []string{now}
		}

		for i, to := range graph[now] {
			if to == "" {
				continue
			}
			graph[now][i] = ""
			result := helper(to, n-1)
			if len(result) != 0 {
				return append(append(res, now), result...)
			}
			graph[now][i] = to
		}
		return res
	}

	return helper("JFK", len(tickets))
}
