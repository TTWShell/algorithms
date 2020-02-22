/* https://leetcode.com/problems/scramble-string/
Given a string s1, we may represent it as a binary tree
by partitioning it to two non-empty substrings recursively.

Below is one possible representation of s1 = "great":

    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t

To scramble the string, we may choose any non-leaf node
and swap its two children.

For example, if we choose the node "gr" and swap its two children,
it produces a scrambled string "rgeat".

    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
We say that "rgeat" is a scrambled string of "great".

Similarly, if we continue to swap the children of nodes "eat" and "at",
it produces a scrambled string "rgtae".

    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
	  t   a

We say that "rgtae" is a scrambled string of "great".

Given two strings s1 and s2 of the same length,
determine if s2 is a scrambled string of s1.

Example 1:

	Input: s1 = "great", s2 = "rgeat"
	Output: true

Example 2:

	Input: s1 = "abcde", s2 = "caebd"
	Output: false
*/

package ldp

// 思路2：DP
//		既然已经可以递归 + 记忆化了，那么可以想着用 DP 实现
//		考虑我们记忆化其实记录 s1 和 s2 长度相同的子串是否满意题意
//		（因为每次只能将一个串分成两部分，两部分内部可以继续处理，所以必定是子串和子串相比）
//		设 dp[i][j][k] 表示 s1[i:i+k] 和 s2[j:j+k] 是否满足题意
//		初始化： dp[i][j][1] = s1[i] == s2[j]
//		状态转移方程：
//		1. 不交换 s1 当前子串左右两部分
//			只要存在一个 l (0 < l < k) 使得：dp[i][j][l] 和 dp[i+l][j+l][k-l] 都为 true
//			那么 dp[i][j][k] = true ，否则 dp[i][j][k] = false
//		2. 交换 s1 当前子串左右两部分
//			只要存在一个 l (0 < l < k) 使得：dp[i+l][j][k-l] 和 dp[i][j+k-l][l] 都为 true
//			那么 dp[i][j][k] = true ，否则 dp[i][j][k] = false
//
//		还是感觉 DP 想不到就很难，但是只要有一点思路，努力往 DP 想一想，还是很简单都
//		时间复杂度： O(len(s1)^4) ，空间复杂度： O(len(s1)^3)

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	length := len(s1)
	dp := make([][][]bool, length, length)
	for i := 0; i < length; i++ {
		dp[i] = make([][]bool, length, length)
		for j := 0; j < length; j++ {
			dp[i][j] = make([]bool, length+1, length+1)
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	for k := 2; k <= length; k++ { // 长度放在最外层，因为大长度依赖小长度
		for i := 0; i+k <= length; i++ {
			for j := 0; j+k <= length; j++ {
				for l := 1; l < k; l++ {
					if dp[i][j][l] && dp[i+l][j+l][k-l] {
						dp[i][j][k] = true
						break
					}
					if dp[i+l][j][k-l] && dp[i][j+k-l][l] {
						dp[i][j][k] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][length]
}

// func isScramble(s1 string, s2 string) bool { 0ms
// 	length1, length2 := len(s1), len(s2)
// 	if length1 != length2 {
// 		return false
// 	}

// 	if s1 == s2 {
// 		return true
// 	}

// 	count := make([]int, 26)
// 	for idx, ch1 := range s1 {
// 		count[ch1-'a']++
// 		count[s2[idx]-'a']--
// 	}
// 	for _, cnt := range count {
// 		if cnt != 0 {
// 			return false
// 		}
// 	}

// 	for i := 1; i < len(s1); i++ {
// 		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
// 			return true
// 		}
// 		if isScramble(s1[:i], s2[length2-i:]) && isScramble(s1[i:], s2[:length2-i]) {
// 			return true
// 		}
// 	}
// 	return false
// }
