/* https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.
*/

package leetcode

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	length := len(postorder)
	root := &TreeNode{Val: postorder[length-1]}
	for i, val := range inorder {
		if val == root.Val {
			root.Left = buildTree2(inorder[:i], postorder[:i])
			root.Right = buildTree2(inorder[i+1:], postorder[i:length-1])
			break
		}
	}
	return root
}
