/* https://leetcode.com/problems/gas-station/description/
There are N gas stations along a circular route, where the amount of gas at station i is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to its next station (i+1). You begin the journey with an empty tank at one of the gas stations.

Return the starting gas station's index if you can travel around the circuit once, otherwise return -1.

Note:
The solution is guaranteed to be unique.
*/

package lgreedy

func canCompleteCircuit(gas []int, cost []int) int {
	start, total, tank := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			start, total, tank = i+1, total+tank, 0
		}
	}
	total += tank
	if total < 0 {
		return -1
	}
	return start
}
