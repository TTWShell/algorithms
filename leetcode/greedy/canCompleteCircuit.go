/* https://leetcode.com/problems/gas-station/description/
There are N gas stations along a circular route, where the amount of gas at station i is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to its next station (i+1). You begin the journey with an empty tank at one of the gas stations.

Return the starting gas station's index if you can travel around the circuit once, otherwise return -1.

Note:
The solution is guaranteed to be unique.
*/

package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		tank, j := 0, i
		for ; j < len(gas); j++ {
			tank += gas[j] - cost[j]
			if tank < 0 {
				break
			}
		}
		if j != len(gas) {
			continue
		}
		for j := 0; j < i; j++ {
			tank += gas[j] - cost[j]
			if tank < 0 {
				break
			}
		}
		if tank >= 0 {
			return i
		}
	}
	return -1
}
