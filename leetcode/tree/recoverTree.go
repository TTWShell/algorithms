/* https://leetcode.com/problems/recover-binary-search-tree/
Two elements of a binary search tree (BST) are swapped by mistake.

Recover the tree without changing its structure.

Example 1:

Input: [1,3,null,null,2]

   1
  /
 3
  \
   2

Output: [3,1,null,null,2]

   3
  /
 1
  \
   2
Example 2:

Input: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

Output: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3
Follow up:

	A solution using O(n) space is pretty straight forward.
	Could you devise a constant space solution?
*/

package ltree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	// 中序遍历bst有序。每次记住前一个节点，寻找到第一个乱序的必定为大的那个
	// 接下来找到的节点本身为小的那个
	// 例如 1 2 3 4 5 6 --> 1 5 3 4 2 6
	var pre, first, second *TreeNode

	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inOrder(root.Left)

		if pre != nil && pre.Val > root.Val {
			if first == nil {
				first = pre
			}
			second = root
		}
		pre = root

		inOrder(root.Right)
	}

	inOrder(root)
	first.Val, second.Val = second.Val, first.Val
}
