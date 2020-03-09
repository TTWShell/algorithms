/* https://leetcode.com/problems/dungeon-game/
The demons had captured the princess (P) and imprisoned her in the
bottom-right corner of a dungeon.
The dungeon consists of M x N rooms laid out in a 2D grid.
Our valiant knight (K) was initially positioned in the top-left room
and must fight his way through the dungeon to rescue the princess.

The knight has an initial health point represented by a positive integer.
 If at any point his health point drops to 0 or below,
 he dies immediately.

Some of the rooms are guarded by demons, so the knight loses health
(negative integers) upon entering these rooms; other rooms are either
empty (0's) or contain magic orbs that increase the knight's
health (positive integers).

In order to reach the princess as quickly as possible,
the knight decides to move only rightward or downward in each step.



Write a function to determine the knight's minimum initial health
so that he is able to rescue the princess.

For example, given the dungeon below, the initial health of the knight
must be at least 7 if he follows the optimal
path RIGHT-> RIGHT -> DOWN -> DOWN.

-2 (K)	-3	3
-5		-10	1
10		30	-5 (P)


Note:

The knight's health has no upper bound.
Any room can contain threats or power-ups, even the first room the
 knight enters and the bottom-right room where the princess is imprisoned.
*/

package ldp

import (
	"math"
)

func calculateMinimumHP(dungeon [][]int) int {
	// 正序处理无法解决下个节点依赖问题，逆推
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	m, n := len(dungeon), len(dungeon[0])
	dp := make([]int, n+1, n+1) // 进入房间必须的最小hp，优化：不需要缓存之前的行
	// 初始化边界
	for i := 0; i <= n; i++ {
		dp[i] = math.MaxInt64
	}

	dp[n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 当前房间取决于之后进入的房间所需最小值, dp[j]保存的为上一行的结果
			dp[j] = max(1, min(dp[j], dp[j+1])-dungeon[i][j])
		}
	}
	return dp[0]
}
