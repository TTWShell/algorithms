/* https://leetcode.com/problems/judge-route-circle/description/
Initially, there is a Robot at position (0, 0). Given a sequence of its moves, judge if this robot makes a circle, which means it moves back to the original place.

The move sequence is represented by a string. And each move is represent by a character. The valid robot moves are R (Right), L (Left), U (Up) and D (down).
The output should be true or false representing whether the robot makes a circle.

Example 1:
    Input: "UD"
    Output: true

Example 2:
    Input: "LL"
    Output: false
*/
package leetcode

func judgeCircle(moves string) bool {
	directions := make([]int, 4, 4)
	for _, direction := range moves {
		switch direction {
		case 'R':
			directions[0]++
		case 'L':
			directions[1]++
		case 'U':
			directions[2]++
		case 'D':
			directions[3]++
		}
	}
	return directions[0] == directions[1] && directions[2] == directions[3]
}
