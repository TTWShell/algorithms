package tree

// PostOrderRecursion : 后序遍历的递归实现。
func PostOrderRecursion(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	res = append(res, PostOrderRecursion(root.Left)...)
	res = append(res, PostOrderRecursion(root.Right)...)
	res = append(res, root.Val)
	return res
}
