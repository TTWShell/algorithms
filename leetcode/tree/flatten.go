/* https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/
Given a binary tree, flatten it to a linked list in-place.

For example,
    Given

             1
            / \
           2   5
          / \   \
         3   4   6
    The flattened tree should look like:
       1
        \
         2
          \
           3
            \
             4
              \
               5
                \
                 6
Hints:
If you notice carefully in the flattened tree, each node's right child points to the next node of a pre-order traversal.
*/

package leetcode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var helper func(root *TreeNode) *TreeNode
	helper = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		var pLeft, left, right *TreeNode
		res := &TreeNode{Val: root.Val}
		left, right = helper(root.Left), helper(root.Right)
		curLeft := left

		for curLeft != nil {
			pLeft, curLeft = curLeft, curLeft.Right
		}
		if pLeft != nil {
			res.Right = left
			pLeft.Right = right
		} else {
			res.Right = right
		}
		return res
	}

	root.Right = helper(root).Right
	root.Left = nil
}
