/* https://leetcode.com/problems/divisor-game/
Alice and Bob take turns playing a game, with Alice starting first.

Initially, there is a number N on the chalkboard.
On each player's turn, that player makes a move consisting of:

	Choosing any x with 0 < x < N and N % x == 0.
	Replacing the number N on the chalkboard with N - x.

Also, if a player cannot make a move, they lose the game.

Return True if and only if Alice wins the game, assuming both players play optimally.


Example 1:

	Input: 2
	Output: true
	Explanation: Alice chooses 1, and Bob has no more moves.

Example 2:

	Input: 3
	Output: false
	Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.

Note:

1 <= N <= 1000
*/

package ldp

import "math"

func divisorGame(N int) bool {
	if N == 1 {
		return false
	}

	dp := make([]int, N+1, N+1) // 0 unknown, -1 false, 1 true
	dp[1], dp[2] = -1, 1
	for i := 3; i <= N; i++ {
		for j := 1; j <= int(math.Sqrt(float64(N))); j++ {
			if tmp := i - j; i%j == 0 && dp[tmp] != 0 {
				dp[i] = -dp[tmp]
				if dp[i] == 1 {
					break
				}
			}
		}
	}
	return dp[N] == 1
}
