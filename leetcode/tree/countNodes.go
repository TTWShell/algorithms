/* https://leetcode.com/problems/count-complete-tree-nodes/description/
Given a complete binary tree, count the number of nodes.

Note:

Definition of a complete binary tree from Wikipedia:
In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

Example:

    Input:
        1
       / \
      2   3
     / \  /
    4  5 6

    Output: 6
*/

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res, lastLevelCount := 0, 1
	for tmp := root; tmp != nil; tmp = tmp.Left {
		if res != 0 {
			lastLevelCount *= 2
		}
		res += lastLevelCount
	}

	start, end := 1, lastLevelCount
	for start <= end {
		cur, mid, prefix := root, start+(end-start)>>1, 0
		for length := lastLevelCount; length > 1; length /= 2 {
			if tmp := length / 2; mid <= prefix+tmp {
				cur = cur.Left
			} else {
				cur = cur.Right
				prefix += tmp
			}
		}
		if cur != nil {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return res - (lastLevelCount - end)
}
