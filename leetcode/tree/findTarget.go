/* https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/
Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:
    Input:
        5
       / \
      3   6
     / \   \
    2   4   7

    Target = 9

    Output: True
Example 2:
    Input:
        5
       / \
      3   6
     / \   \
    2   4   7

    Target = 28

    Output: False
*/

package ltree

func findTarget(root *TreeNode, k int) bool {
	return find2(root, root, k)
}

func find2(root *TreeNode, curnode *TreeNode, k int) bool {
	if curnode == nil {
		return false
	}
	if k != curnode.Val*2 && findNode(root, k-curnode.Val) {
		return true
	}
	return find2(root, curnode.Left, k) || find2(root, curnode.Right, k)
}

func findNode(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}

	if root.Val == val {
		return true
	} else if root.Val < val {
		return findNode(root.Right, val)
	}
	return findNode(root.Left, val)
}
