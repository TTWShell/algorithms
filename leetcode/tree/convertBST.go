/* https://leetcode.com/problems/convert-bst-to-greater-tree/description/
Given a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.

Example:

    Input: The root of a Binary Search Tree like this:
                  5
                /   \
               2     13

    Output: The root of a Greater Tree like this:
                 18
                /   \
              20     13
*/

package ltree

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var helper func(root *TreeNode) (res *TreeNode)
	helper = func(root *TreeNode) (res *TreeNode) {
		if root != nil {
			res = &TreeNode{Val: root.Val}
			res.Right = helper(root.Right)
			res.Val += sum
			sum = res.Val
			res.Left = helper(root.Left)
		}
		return res
	}
	return helper(root)
}
