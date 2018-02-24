/* https://leetcode.com/problems/minimum-distance-between-bst-nodes/description/
Given a Binary Search Tree (BST) with the root node root, return the minimum difference between the values of any two different nodes in the tree.

Example :

Input: root = [4,2,6,1,3,null,null]
Output: 1
Explanation:
Note that root is a TreeNode object, not an array.

The given tree [4,2,6,1,3,null,null] is represented by the following diagram:

          4
        /   \
      2      6
     / \
    1   3

while the minimum difference in this tree is 1, it occurs between node 1 and node 2, also between node 3 and node 2.
Note:

The size of the BST will be between 2 and 100.
The BST is always valid, each node's value is an integer, and each node's value is different.
*/

package leetcode

func minDiffInBST(root *TreeNode) int {
	pre, cur, res := root.Val, root.Val, 0
	stack := []*TreeNode{}

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		pre, cur = cur, top.Val
		if tmp := cur - pre; tmp > 0 && (tmp < res || res == 0) {
			res = tmp
		}
		root = top.Right

	}
	return res
}
